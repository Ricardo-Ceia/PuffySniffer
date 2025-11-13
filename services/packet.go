package services

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func ListAllDevs() {
	interfaces, err := pcap.FindAllDevs()
	if err != nil {
		panic(err)
	}
	for i := range interfaces {
		log.Println(interfaces[i])
	}
}

func OpenLiveToAny() *pcap.Handle {
	handle, err := pcap.OpenLive("eth0", 65535, true, pcap.BlockForever)

	if err != nil {
		panic(err)
	}

	return handle
}

func IteratePackets(h *pcap.Handle) {
	packetSource := gopacket.NewPacketSource(h, h.LinkType())

	for packet := range packetSource.Packets() {
		log.Println(packet)
	}
}
