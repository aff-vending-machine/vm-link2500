package main

import (
	"fmt"
	"log"
	"time"

	"github.com/tarm/serial"
)

func main() {
	config := &serial.Config{
		Name:        "/dev/ttyACM0",
		Baud:        9600,
		ReadTimeout: 5 * time.Minute,
		Size:        8,
	}

	stream, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	err = stream.Flush()
	if err != nil {
		log.Fatal(err)
	}

	sale(stream) // tap card 3 steps
	time.Sleep(5 * time.Second)
	void(stream) // auto  3 steps
	time.Sleep(5 * time.Second)
	refund(stream) // tap card  3 steps
	time.Sleep(5 * time.Second)
	sale(stream) // tap card  3 steps
	time.Sleep(5 * time.Second)
	settlement(stream) // auto 2 steps
	time.Sleep(5 * time.Second)
}

func sale(stream *serial.Port) {
	fmt.Println("Sale")
	b := []byte{0x02, 0x00, 0x46, 0x36, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x32, 0x30, 0x30, 0x30, 0x30, 0x1C, 0x34, 0x30, 0x00, 0x12, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x32, 0x35, 0x31, 0x35, 0x1C, 0x34, 0x35, 0x00, 0x06, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x1C, 0x03, 0x7B}
	n, err := stream.Write(b)
	if err != nil {
		log.Fatalf("stream.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")

	resp := make([]byte, 1)
	n, err = stream.Read(resp)
	if err != nil {
		log.Fatalf("stream.Read: %v", err)
	}

	if n == 1 && resp[0] == 0x06 {
		fmt.Printf("Read %d bytes. %v\n", n, resp[:n])
	} else {
		log.Fatalf("Read: %v", resp[:n])
	}

	buf := make([]byte, 256)
	n, err = stream.Read(buf)
	if err != nil {
		log.Fatalf("stream.Read: %v", err)
	}

	if n > 0 {
		fmt.Printf("Read %d bytes. %v\n", n, buf[:n])
	}
}

func void(stream *serial.Port) {
	fmt.Println("Void")
	b := []byte{0x02, 0x00, 0x29, 0x36, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x32, 0x36, 0x30, 0x30, 0x30, 0x1C, 0x36, 0x35, 0x00, 0x06, 0x20, 0x20, 0x20, 0x20, 0x20, 0x31, 0x1C, 0x03, 0x0B}
	n, err := stream.Write(b)
	if err != nil {
		log.Fatalf("stream.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")

	resp := make([]byte, 1)
	n, err = stream.Read(resp)
	if err != nil {
		log.Fatalf("stream.Read: %v", err)
	}

	if n == 1 && resp[0] == 0x06 {
		fmt.Printf("Read %d bytes. %v\n", n, resp[:n])
	} else {
		log.Fatalf("Read: %v", resp[:n])
	}

	buf := make([]byte, 256)
	n, err = stream.Read(buf)
	if err != nil {
		log.Fatalf("stream.Read: %v", err)
	}

	if n > 0 {
		fmt.Printf("Read %d bytes. %v\n", n, buf[:n])
	}
}

func refund(stream *serial.Port) {
	fmt.Println("Refund")
	b := []byte{0x02, 0x00, 0x46, 0x36, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x32, 0x37, 0x30, 0x30, 0x30, 0x1C, 0x34, 0x30, 0x00, 0x12, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x32, 0x35, 0x32, 0x35, 0x1C, 0x34, 0x35, 0x00, 0x06, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x1C, 0x03, 0x7C}
	n, err := stream.Write(b)
	if err != nil {
		log.Fatalf("stream.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")

	resp := make([]byte, 1)
	n, err = stream.Read(resp)
	if err != nil {
		log.Fatalf("stream.Read: %v", err)
	}

	if n == 1 && resp[0] == 0x06 {
		fmt.Printf("Read %d bytes. %v\n", n, resp[:n])
	} else {
		log.Fatalf("Read: %v", resp[:n])
	}

	buf := make([]byte, 256)
	n, err = stream.Read(buf)
	if err != nil {
		log.Fatalf("stream.Read: %v", err)
	}

	if n > 0 {
		fmt.Printf("Read %d bytes. %v\n", n, buf[:n])
	}
}

func settlement(stream *serial.Port) {
	fmt.Println("Settlement")
	b := []byte{0x02, 0x00, 0x37, 0x36, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x30, 0x35, 0x30, 0x30, 0x30, 0x30, 0x1C, 0x48, 0x4E, 0x00, 0x03, 0x39, 0x39, 0x39, 0x1C, 0x34, 0x35, 0x00, 0x06, 0x30, 0x30, 0x30, 0x30, 0x30, 0x31, 0x1C, 0x03, 0x20}
	n, err := stream.Write(b)
	if err != nil {
		log.Fatalf("stream.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")

	resp := make([]byte, 1)
	n, err = stream.Read(resp)
	if err != nil {
		log.Fatalf("stream.Read: %v", err)
	}

	if n == 1 && resp[0] == 0x06 {
		fmt.Printf("Read %d bytes. %v\n", n, resp[:n])
	} else {
		log.Fatalf("Read: %v", resp[:n])
	}
}
