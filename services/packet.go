package services

import (
	"github.com/google/gopacket/pcap"
	"log"
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
	handle, err := pcap.OpenLive("any", 65535, true, pcap.BlockForever)

	if err != nil {
		panic(err)
	}

	return handle
}
