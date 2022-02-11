package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

type processor func(srcIp string, payload []byte)

func getSrcIp(pkt gopacket.Packet) string {
	ipLayer := pkt.NetworkLayer()
	if ipLayer == nil {
		return ""
	}

	// ipv4 or ipv6
	switch l := ipLayer.(type) {
	case *layers.IPv4:
		return l.SrcIP.String()
	case *layers.IPv6:
		return l.SrcIP.String()
	}
	return ""
}

func getUdpPayload(pkt gopacket.Packet) []byte {
	layer := pkt.TransportLayer()
	if layer == nil {
		return nil
	}

	if layer.LayerType() != layers.LayerTypeUDP {
		return nil
	}

	switch l := layer.(type) {
	case *layers.UDP:
		return l.Payload
	}
	return nil
}

func PacketLoop(dev string, udpPort int, p processor) error {
	handler, err := pcap.OpenLive(dev, 1496, false, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handler.Close()

	if err := handler.SetBPFFilter(fmt.Sprintf("udp and port %d", udpPort)); err != nil {
		return err
	}

	source := gopacket.NewPacketSource(handler, handler.LinkType())
	for pkt := range source.Packets() {
		srcIp := getSrcIp(pkt)
		payload := getUdpPayload(pkt)

		if srcIp != "" && payload != nil {
			p(srcIp, payload)
		}
	}

	return nil
}
