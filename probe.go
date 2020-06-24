package main

import (
	"github.com/Team-SYNACKd/probe/pkg/tuntap"
	"github.com/songgao/packets/ethernet"
)

func main() {
	var frame ethernet.Frame
	tuntap.TunRead(frame)

}
