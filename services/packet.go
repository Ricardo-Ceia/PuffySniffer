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
