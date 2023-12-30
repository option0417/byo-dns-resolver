package main

import (
	"fmt"
	"testing"
)

func TestToBytes(t *testing.T) {
	header := Header{}

	var v uint16
	for v = 1; v <= 32768; v *= 2 {
		fmt.Printf("Value: %d\n", v)
		header.Id = v
		bytes, _ := header.ToBytes()
		show(bytes)

		if v == 32768 {
			break
		}
	}
}

func show(bytes []byte) {
	fmt.Printf("Length: %d\n", len(bytes))
	fmt.Printf("%08b %08b\n", bytes[0], bytes[1])
}
