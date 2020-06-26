package physical

import (
	"log"

	"github.com/songgao/packets/ethernet"
	"github.com/songgao/water"
)

func PrepareTunInterface() *water.Interface {
	config := water.Config{
		DeviceType: water.TUN,
	}
	config.Name = "probe"

	ifce, err := water.New(config)
	if err != nil {
		log.Fatal(err)
	}

	return ifce
}

func TunRead(ifce *water.Interface, frame ethernet.Frame) int {
	n, err := ifce.Read([]byte(frame))
	if err != nil {
		log.Fatal(err)
	}

	return n
}
