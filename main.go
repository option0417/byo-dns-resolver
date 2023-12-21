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
)

func main() {
	queryString := "00160100000100000000000003646e7306676f6f676c6503636f6d0000010001"

	queryBytes, err := hex.DecodeString(queryString)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("udp", DNS_GoogleIP)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	conn.SetReadDeadline(time.Now().Add(time.Second * 3))

	cnt, err := conn.Write(queryBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Write %d bytes\n", cnt)

	buf := make([]byte, 100)
	readCnt := -1
	var readErr error
	cnt = 0

	for readCnt != 0 {
		fmt.Printf("New\n")
		cnt++
		readCnt, readErr = conn.Read(buf)
		if readErr != nil {
			panic(readErr)
		}
		fmt.Printf("Read %d bytes\n", readCnt)
		fmt.Printf("Round: %d\n", cnt)

		buf = make([]byte, 100)
	}

}
