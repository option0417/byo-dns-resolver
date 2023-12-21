package main

import (
	"encoding/hex"
	"fmt"
)

const (
	DNS_Google   = "dns.google.com"
	DNS_GoogleIP = "8.8.8.8"
)

func main() {
	queryString := "00160100000100000000000003646e7306676f6f676c6503636f6d0000010001"

	queryBytes, err := hex.DecodeString(queryString)
	if err != nil {
		panic(err)
	}

	fmt.Println("vim-go")
}
