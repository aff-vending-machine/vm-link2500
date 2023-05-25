package link2500

import (
	"context"
	"fmt"
	"strings"
	"time"

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
func (e *serialImpl) Sale(ctx context.Context, req *request.Sale) (*response.Result, error) {
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

	price := fmt.Sprintf("%012d", int(req.Price*100))

	stx := []byte{0x02}
	etx := []byte{0x03}
	th := makeTransportHeader()
	ph := makePresentationHeader("0", "20")
	fd1 := makeFieldData("40", []byte(price))
	fd2 := makeFieldData("45", []byte(req.MerchantID))
	msg := concat(th, ph, fd1, fd2)
	llll := get2ByteLength(msg)
	payload := concat(llll, msg, etx)
	lrc := calLRC(payload)
	payload = concat(stx, payload, []byte{lrc})

	// 1. POS send request to EDC
	log.Info().Str("payload", toHex(payload)).Msg("EDC: (1) send")
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

	log.Info().Int("length", n).Str("result", toHex(result1[:n])).Msg("EDC: (2) received")
	if n != 1 || result1[0] != 0x06 {
		return nil, fmt.Errorf("receive unknown message (%d): %v", n, result1[:n])
	}

	// 3. EDC return Result to POS
	result2 := make([]byte, 1024)
	n, err = stream.Read(result2)
	if err != nil {
		log.Warn().Err(err).Msg("EDC: (3) received error, need to manual inquiry")

		count := 0
		for {
			select {
			case <-ctx.Done():
				log.Warn().Msg("inquiry cancelled")
				return nil, fmt.Errorf("cancelled")

			default:
				count++
				if count == 100 {
					log.Warn().Msg("inquiry cancelled")
					return nil, fmt.Errorf("cancelled by system")
				}

				log.Info().Msg("inquiry")
				result, err := e.inquiry(stream)
				if err != nil {
					return nil, err
				}

				if len(result) == 1 {
					log.Err(err).Str("result", toHex(result)).Msg("response data delay retry")
					stream.Flush()
					time.Sleep(3 * time.Second)
					continue
				}

				if result[0] == 0x06 && result[1] == 0x02 {
					result = result[1:]
				}

				if result[0] != 0x02 {
					log.Err(err).Str("result", toHex(result)).Msg("noise occured, need to flush data and re-inquiry")
					stream.Flush()
					time.Sleep(3 * time.Second)
					continue
				}

				if len(result) < 24 {
					log.Err(err).Str("result", toHex(result)).Msg("response data is incorrectly")
					stream.Flush()
					time.Sleep(3 * time.Second)
					continue
				}

				edcInquiry := generateResult(result)

				if !strings.Contains(edcInquiry.ResponseText, "APPROVED") {
					time.Sleep(3 * time.Second)
					continue
				}

				return edcInquiry, nil
			}
		}

	} else {
		result := result2[:n]
		log.Info().Int("length", n).Str("result", toHex(result)).Msg("EDC (3) received")

		edcResult := generateResult(result)
		log.Info().Interface("result", edcResult).Msg("EDC (3) received")

		// 4. POS response ACK to EDC
		log.Info().Str("payload", "06").Msg("EDC: (4) send")
		_, err = stream.Write([]byte{0x06})
		if err != nil {
			return nil, err
		}

		return edcResult, nil
	}
}

func (*serialImpl) inquiry(stream *serial.Port) ([]byte, error) {
	stx := []byte{0x02}
	etx := []byte{0x03}
	th := makeTransportHeader()
	ph := makePresentationHeader("0", "74")
	msg := concat(th, ph)
	llll := get2ByteLength(msg)
	payload := concat(llll, msg, etx)
	lrc := calLRC(payload)
	payload = concat(stx, payload, []byte{lrc})

	log.Info().Str("payload", toHex(payload)).Msg("EDC (3.1) send inquiry")
	_, err := stream.Write(payload)
	if err != nil {
		return nil, err
	}

	result := make([]byte, 2048)
	n, err := stream.Read(result)
	if err != nil {
		return nil, err
	}
	log.Info().Int("length", n).Str("result", toHex(result[:n])).Msg("EDC (3.2) received inquiry")

	return result[:n], nil
}
