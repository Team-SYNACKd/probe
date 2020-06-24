package v4

import (
	"log"
	"net"

	"github.com/songgao/packets/ethernet"
	"golang.org/x/net/ipv4"
)

type Ipv4Header struct {
	SrcAddr    net.IP
	DestAddr   net.IP
	TotalLen   int
	HeaderLen  int
	PayloadLen int
	Protocol   int
}

func Parsev4(frame ethernet.Frame) Ipv4Header {

	header, err := ipv4.ParseHeader(frame)
	if err != nil {
		log.Fatal(err)
	}

	hdr := Ipv4Header{
		SrcAddr:    header.Src,
		DestAddr:   header.Dst,
		TotalLen:   header.TotalLen,
		HeaderLen:  header.Len,
		PayloadLen: header.TotalLen - header.Len,
		Protocol:   header.Protocol,
	}

	return hdr
}
