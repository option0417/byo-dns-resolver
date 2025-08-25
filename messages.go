package main

import "strings"

// DNS Message Format
//  +---------------------+
//  |        Header       |
//  +---------------------+
//  |       Question      | the question for the name server
//  +---------------------+
//  |        Answer       | RRs answering the question
//  +---------------------+
//  |      Authority      | RRs pointing toward an authority
//  +---------------------+
//  |      Additional     | RRs holding additional information
//  +---------------------+

// Header section
//
//	                              1  1  1  1  1  1
//	0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                      ID                       |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|QR|   Opcode  |AA|TC|RD|RA|   Z    |   RCODE   |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                    QDCOUNT                    |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                    ANCOUNT                    |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                    NSCOUNT                    |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                    ARCOUNT                    |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
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

func (this *Header) fetchIDBytes() []byte {
	if (this.Id == 0) || (this.Id > 65535) {
		this.Id = uint16(1 + (randInt() % 65534))
	}

	return []byte{
		byte((this.Id >> 8) & 0xFF),
		byte(this.Id & 0xFF),
	}
}

func (this *Header) fetchFlagBytes() []byte {
	var flags uint16 = 0

	if this.QR {
		flags |= 1 << 15
	}

	flags |= (uint16(this.OPcode) & 0x0F) << 11

	if this.AA {
		flags |= 1 << 10
	}

	if this.TC {
		flags |= 1 << 9
	}

	if this.RD {
		flags |= 1 << 8
	}

	if this.RA {
		flags |= 1 << 7
	}

	// Z is reserved and must be zero

	flags |= uint16(this.RCode) & 0x0F

	return []byte{
		byte((flags >> 8) & 0xFF),
		byte(flags & 0xFF),
	}
}

func (this *Header) fetchQDCountBytes() []byte {
	return []byte{
		byte((this.QDCount >> 8) & 0xFF),
		byte(this.QDCount & 0xFF),
	}
}

func (this *Header) fetchANCountBytes() []byte {
	return []byte{
		byte((this.ANCount >> 8) & 0xFF),
		byte(this.ANCount & 0xFF),
	}
}

func (this *Header) fetchNSCountBytes() []byte {
	return []byte{
		byte((this.NSCount >> 8) & 0xFF),
		byte(this.NSCount & 0xFF),
	}
}

func (this *Header) fetchARCountBytes() []byte {
	return []byte{
		byte((this.ARCount >> 8) & 0xFF),
		byte(this.ARCount & 0xFF),
	}
}

func (this *Header) toBytes() []byte {
	bytes := []byte{}
	bytes = append(bytes, this.fetchIDBytes()...)
	bytes = append(bytes, this.fetchFlagBytes()...)
	bytes = append(bytes, this.fetchQDCountBytes()...)
	bytes = append(bytes, this.fetchANCountBytes()...)
	bytes = append(bytes, this.fetchNSCountBytes()...)
	bytes = append(bytes, this.fetchARCountBytes()...)

	return bytes
}

// Question section
//
//	                                1  1  1  1  1  1
//	  0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                                               |
//	/                     QNAME                     /
//	/                                               /
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                     QTYPE                     |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                     QCLASS                    |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
type Question struct {
	QName  string
	QType  uint16
	QClass uint16
}

func (this *Question) fetchQNameBytes() []byte {
	qNameParts := strings.Split(this.QName, ".")

	var qNameBytes []byte

	for _, part := range qNameParts {
		length := len(part)
		if length > 0 && length < 64 {
			qNameBytes = append(qNameBytes, byte(length))
			qNameBytes = append(qNameBytes, []byte(part)...)
		}
	}

	qNameBytes = append(qNameBytes, 0) // End of QNAME
	return qNameBytes
}

func (this *Question) fetchQTypeBytes() []byte {
	return []byte{
		byte((this.QType >> 8) & 0xFF),
		byte(this.QType & 0xFF),
	}
}

func (this *Question) fetchQClassBytes() []byte {
	return []byte{
		byte((this.QClass >> 8) & 0xFF),
		byte(this.QClass & 0xFF),
	}
}

// Resource record format
//
//	                                1  1  1  1  1  1
//	  0  1  2  3  4  5  6  7  8  9  0  1  2  3  4  5
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                                               |
//	/                                               /
//	/                      NAME                     /
//	|                                               |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                      TYPE                     |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                     CLASS                     |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                      TTL                      |
//	|                                               |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//	|                   RDLENGTH                    |
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--|
//	/                     RDATA                     /
//	/                                               /
//	+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
type ResourceRecoed struct {
	Name     string
	Type     uint16
	Class    uint16
	TTL      uint16
	RDLength uint16
	RData    uint16
}

// Message compression
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
//    | 1  1|                OFFSET                   |
//    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
