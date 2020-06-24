package main

import (
	"log"

	"github.com/Team-SYNACKd/probe/pkg/tuntap"
	"github.com/songgao/packets/ethernet"
)

func receivePackets() {
	var frame ethernet.Frame
	frame.Resize(1500)

	ifce := tuntap.PrepareTunInterface()

	for {
		n := tuntap.TunRead(ifce, frame)
		frame = frame[:n]
		log.Printf("Dst: %s\n", frame.Destination())
		log.Printf("Src: %s\n", frame.Source())
		log.Printf("Ethertype: % x\n", frame.Ethertype())
		log.Printf("Payload: % x\n", frame.Payload())
	}
}

func main() {
	receivePackets()
}
