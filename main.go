package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"time"
)

const (
	DNS_Google   = "dns.google.com:53"
	DNS_GoogleIP = "8.8.8.8:53"
	DNS_MitakeIP = "10.3.1.254:53"
)

// Header section
//                                  1  1  1  1  1  1
//    0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                      ID                       |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                    QDCOUNT                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                    ANCOUNT                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                    NSCOUNT                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    |                    ARCOUNT                    |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
type Header struct {
	Id      uint16
	QR      bool
	OPcode  uint8
	AA      bool
	TC      bool
	RD      bool
	RA      bool
	RCode   uint8
	QDCount uint16
	ANCount uint16
	NSCount uint16
	ARCount uint16
}

func main() {
	queryString := "00160100000100000000000003646e7306676f6f676c6503636f6d0000010001"

	queryBytes, err := hex.DecodeString(queryString)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", queryString)
	for _, v := range queryBytes {
		fmt.Printf(" %d", v)
	}
	fmt.Println()

	conn, err := net.Dial("udp", DNS_MitakeIP)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(time.Second * 30))

	cnt, err := conn.Write(queryBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Write %d bytes\n", cnt)

	buf := make([]byte, 100)
	readCnt, readErr := conn.Read(buf)
	if readErr != nil {
		panic(readErr)
	}
	fmt.Printf("Read %d bytes\n", readCnt)

	result := hex.EncodeToString(buf[:readCnt])
	fmt.Printf("Result: %s\n", result)
	for _, v2 := range buf {
		fmt.Printf(" %d", v2)
	}
	fmt.Println()
}
