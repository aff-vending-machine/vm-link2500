package link2500

import (
	"context"
	"fmt"

	"vm-link2500/internal/core/infra/peripheral/serial"
	"vm-link2500/internal/layer/usecase/link2500/request"
	"vm-link2500/internal/layer/usecase/link2500/response"

	"github.com/rs/zerolog/log"
)

func (e *serialImpl) Void(ctx context.Context, req *request.Void) (*response.Result, error) {
	stream, err := serial.OpenPort(e.config)
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	if err := stream.Flush(); err != nil {
		return nil, err
	}

	invoice := fmt.Sprintf("%06s", req.InvoiceNumber)

	stx := []byte{0x02}
	etx := []byte{0x03}
	th := makeTransportHeader()
	ph := makePresentationHeader("0", "26")
	fd := makeFieldData("65", []byte(invoice))
	msg := concat(th, ph, fd)
	llll := get2ByteLength(msg)
	payload := concat(llll, msg, etx)
	lrc := calLRC(payload)
	payload = concat(stx, payload, []byte{lrc})

	// 1. POS send request to EDC
	log.Info().Str("payload", toHex(payload)).Msg("EDC: (1) send")
	err = stream.Write(ctx, payload)
	if err != nil {
		return nil, err
	}

	// 2. EDC response ACK to POS
	result1 := make([]byte, 1)
	n, err := stream.Read(ctx, result1)
	if err != nil {
		return nil, err
	}

	log.Info().Str("result", toHex(result1[:n])).Msg("EDC: (2) received")
	if n != 1 || result1[0] != 0x06 {
		return nil, fmt.Errorf("receive unknown message (%d): %v", n, result1[:n])
	}

	// 3. EDC return Result to POS
	result2 := make([]byte, 1024)
	n, err = stream.Read(ctx, result2)
	if err != nil {
		return nil, err
	}

	result := result2[:n]
	log.Info().Str("result", toHex(result)).Msg("EDC (3) received")

	edcResult := generateResult(result)
	log.Info().Interface("result", edcResult).Msg("EDC (3) received")

	// 4. POS response ACK to EDC
	log.Info().Str("payload", "06").Msg("EDC: (4) send")
	err = stream.Write(ctx, []byte{0x06})
	if err != nil {
		return nil, err
	}

	return edcResult, nil
}
