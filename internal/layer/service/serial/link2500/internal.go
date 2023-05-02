package link2500

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/aff-vending-machine/vm-link2500/internal/layer/usecase/link2500/response"
)

func makeTransportHeader() []byte {
	return []byte{
		0x36, 0x30, // Transport Header Type
		0x30, 0x30, 0x30, 0x30, // Transport Destination
		0x30, 0x30, 0x30, 0x30, // Transport Source
	}
}

func makePresentationHeader(rri string, code string) []byte {
	return []byte{
		0x31,                         // Format Version
		byte(rri[0]),                 // Request-Response Indicator
		byte(code[0]), byte(code[1]), // Transaction Code
		0x30, 0x30, // Response Code
		0x30, // More Data Indicator
		0x1C, // Field Separator
	}
}

func makeFieldData(ft string, data []byte) []byte {
	return concat([]byte(ft), get2ByteLength(data), data, []byte{0x1C})
}

func get2ByteLength(data []byte) []byte {
	size := len(data)
	//	return []byte{byte(size / 100), byte(size % 100)}
	result, _ := hex.DecodeString(fmt.Sprintf("%02d%02d", size/100, size%100))
	return result
}

func concat(byteArrays ...[]byte) []byte {
	// https://stackoverflow.com/questions/29046963/concat-byte-arrays
	size := 0
	for _, bytes := range byteArrays {
		size = size + len(bytes)
	}

	buf := make([]byte, 0, size)
	buf = buf[:0]

	for _, bytes := range byteArrays {
		buf = append(buf, bytes...)
	}

	return buf
}

func calLRC(bytes []byte) byte {
	var lrc byte = 0
	for _, b := range bytes {
		lrc ^= b
	}

	return lrc
}

func generateResult(buf []byte) *response.Result {
	// exclude stx, llll, th, ph, last 0x1c, etx, lrc
	msg := buf[21 : len(buf)-3]
	fields := bytes.Split(msg, []byte{0x1C})
	result := &response.Result{}

	for _, fd := range fields {
		data := string(fd[4:])
		switch string(fd[:2]) {
		case "01": // 6
			result.ApprovalCode = data

		case "02": // 40
			result.ResponseText = data

		case "03": // 6
			result.TransactionDate = data

		case "04": // 6
			result.TransactionTime = data

		case "16": // 8
			result.TerminalIdentificationNumber = data

		case "30": // 16
			result.PrimaryAccountNumber = data

		case "31": // 4
			result.ExpirationDate = data

		case "40": // 12
			result.Amount = data

		case "45": // 6
			result.MerchantNumber = data

		case "50": // 6
			result.BatchNumber = data

		case "65": // 6
			result.InvoiceNumber = data

		case "D0": // 69
			result.MerchantName = data

		case "D1": // 15
			result.ISO8583MerchantNumber = data

		case "D2": // 10
			result.CardIssuerName = data

		case "D3": // 12
			result.RetrievalReferenceNumber = data

		case "D4": // 2
			result.CardIssuerID = data

		case "D5": // 26
			result.CardHolderName = data
		}
	}

	return result
}

func generateSettlementResult(buf []byte) *response.Settlement {
	// exclude stx, llll, th, ph, last 0x1c, etx, lrc
	msg := buf[21 : len(buf)-3]
	fields := bytes.Split(msg, []byte{0x1C})
	result := &response.Settlement{}

	for _, fd := range fields {
		data := string(fd[4:])
		switch string(fd[:2]) {
		case "ZZ": // 8 * n
			result.Hosts = strings.SplitN(data, "", 8)
		}
	}

	return result
}

func toHex(byteArray []byte) string {
	hexString := ""
	for _, b := range byteArray {
		hexString += fmt.Sprintf("%02x ", b)
	}

	return hexString
}
