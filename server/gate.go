package main

import "fmt"

// SPAProcessor Single-Packet-Authorization Processor
type SPAProcessor struct {
	firewall Firewall
}

// Process 处理流程
func (p *SPAProcessor) Process(srcIp string, payload []byte) {
	fmt.Printf("srcIp: %s payload: %x\n", srcIp, payload)
	if p.Check(payload) {
		p.firewall.Open(srcIp)
	}
}

func (p *SPAProcessor) Check(payload []byte) bool {
	// TODO: 检查SPA包的合法性
	return true
}
