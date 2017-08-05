package main

import (
	"encoding/binary"
	"fmt"

	"github.com/meyskens/uintvar"
)

func getSubPDULength(in []byte) uint16 {
	in[0] &= 0x7F // remove int sign
	return binary.BigEndian.Uint16(in)
}

type wirelessSessionProtocolHeaders struct {
	PDUType            uint16
	MajorVersion       uint16
	MinorVersion       uint16
	CapabilitiesLength uint64
	//Capabilites ignored for now
	HeadersLength uint64
	Headers       []map[uint16]string
}

func parseWSPHeaders(in []byte) (wirelessSessionProtocolHeaders, int) {
	fmt.Println(in)
	out := wirelessSessionProtocolHeaders{}
	c := 0
	out.PDUType = uint16(in[c])
	c++

	version := in[c]
	out.MajorVersion = uint16((version & 0xF0) >> 4) // nill last 4 bytes and move them
	out.MinorVersion = uint16(version & 0xF)
	c++

	var n int
	out.CapabilitiesLength, n, _ = uintvar.Parse(in[c:])
	c += n

	out.HeadersLength, n, _ = uintvar.Parse(in[c:])
	c += n

	//ignore capabilities
	c += int(out.CapabilitiesLength)

	//ignotre till we figure out how indicated
	c += int(out.HeadersLength)

	return out, c
}
