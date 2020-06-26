package main

import (
	"log"

	v4 "github.com/Team-SYNACKd/probe/ip/v4"
	"github.com/Team-SYNACKd/probe/physical"
	"github.com/songgao/packets/ethernet"
)

func receivePackets() {
	var frame ethernet.Frame
	frame.Resize(1500)

	ifce := physical.PrepareTunInterface()

	for {
		n := physical.TunRead(ifce, frame)
		frame = frame[:n]

		//TODO: Should consider only IPv4 frame
		//from frame.Ethertype() until we provide
		//support for v6.
		// if bytes.Compare(frame.Ethertype(), []byte{0x08, 0x00}) != 0 {
		// 	continue
		// }
		//log.Printf("Ethertype %x\n", frame.Ethertype())

		header := v4.Parsev4(frame)

		log.Printf("src: %s dest: %s payload len in bytes: %d protocol %d",
			header.SrcAddr, header.DestAddr, header.PayloadLen, header.Protocol)

		// log.Printf("Dst: %s\n", frame.Destination())
		// log.Printf("Src: %s\n", frame.Source())
		// log.Printf("Ethertype: % x\n", frame.Ethertype())
		// log.Printf("Payload: % x\n", frame.Payload())
	}
}

func main() {
	receivePackets()
}
