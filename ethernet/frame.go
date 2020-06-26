package ethernet

type Ethertype [2]byte

var (
	ARP = Ethertype{0x08, 0x06}
	IPv4 = Ethertype{0x08, 0x00}
	IPv6 = Ethertype{0x86, 0xDD}
)

type EthernetFrame struct {
	Dmac      net.HardwareAddr
	Smac      net.HardwareAddr
	EtherType EtherType
	Payload   []byte
}