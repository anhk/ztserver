package main

import "flag"

var InterfaceName = flag.String("i", "eth0", "interface name")

func main() {

	// Firewall
	firewall := &IptablesFirewall{}
	if err := firewall.Init(*InterfaceName); err != nil {
		panic(err)
	}

	// Single-Packet-Authorization Processor
	spaProcessor := &SPAProcessor{firewall: firewall}

	// PCAP 抓包处理
	if err := PacketLoop(*InterfaceName, 53, spaProcessor.Process); err != nil {
		panic(err)
	}
}
