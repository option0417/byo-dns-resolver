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
	DefaultQuery = "00160100000100000000000003646e7306676f6f676c6503636f6d0000010001"
)

func main() {
	queryBytes, err := hex.DecodeString(DefaultQuery)
	if err != nil {
		panic(err)
	}

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
