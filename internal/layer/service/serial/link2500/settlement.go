package link2500

import (
	"context"
	"fmt"

	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500/request"
	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500/response"
	"github.com/aff-vending-machine/vm-link2500/pkg/trace"
	"github.com/rs/zerolog/log"
	"github.com/tarm/serial"
)

// Process Flow
// 1. POS send request to EDC
// 2. EDC response ACK to POS
// 3. EDC return Result to POS
// 4. POS response ACK to EDC
func (e *serialImpl) Settlement(ctx context.Context, req *request.Settlement) (*response.Settlement, error) {
	_, span := trace.Start(ctx)
	defer span.End()

	stream, err := serial.OpenPort(e.config)
	if err != nil {
		return nil, err
	}
	defer stream.Close()

	if err := stream.Flush(); err != nil {
		return nil, err
	}

	all := "999"

	stx := []byte{0x02}
	etx := []byte{0x03}
	th := makeTransportHeader()
	ph := makePresentationHeader("0", "50")
	fd1 := makeFieldData("HN", []byte(all))
	fd2 := makeFieldData("45", []byte(req.MerchantID))
	msg := concat(th, ph, fd1, fd2)
	llll := get2ByteLength(msg)
	payload := concat(llll, msg, etx)
	lrc := calLRC(payload)
	payload = concat(stx, payload, []byte{lrc})

	// 1. POS send request to EDC
	log.Info().Bytes("payload", payload).Msg("EDC: (1) send")
	_, err = stream.Write(payload)
	if err != nil {
		return nil, err
	}

	// 2. EDC response ACK to POS
	result1 := make([]byte, 1)
	n, err := stream.Read(result1)
	if err != nil {
		return nil, err
	}

	log.Info().Bytes("result", result1[:n]).Msg("EDC: (2) received")
	if n != 1 || result1[0] != 0x06 {
		return nil, fmt.Errorf("receive unknown message (%d): %v", n, result1[:n])
	}

	// 3. EDC return Result to POS
	result2 := make([]byte, 1024)
	n, err = stream.Read(result2)
	if err != nil {
		return nil, err
	}

	result := result2[:n]
	log.Info().Bytes("result", result).Msg("EDC (3) received")
	edcResult := generateSettlementResult(result)

	// 4. POS response ACK to EDC
	log.Info().Bytes("payload", []byte{0x06}).Msg("EDC: (4) send")
	_, err = stream.Write([]byte{0x06})
	if err != nil {
		return nil, err
	}

	return edcResult, nil

}
