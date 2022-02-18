package main

type Firewall interface {
	Init(inf string) error    // 初始化
	Open(srcIp string) error  // 开门
	Close(srcIp string) error // 关门
}

type IptablesFirewall struct {
}

var (
	IpsetName = "ztserver"
)

func (ipt *IptablesFirewall) Init(infName string) error {
	// clear
	_ = RunCmd("iptables", "-D", "INPUT", "-m", "set", "--match-set", IpsetName, "src", "-i", infName, "-j", "ACCEPT")
	_ = RunCmd("ipset", "destroy", IpsetName)

	// init
	_ = RunCmd("ipset", "create", IpsetName, "hash:ip", "timeout", "30")
	_ = RunCmd("iptables", "-I", "INPUT", "1", "-m", "set", "--match-set", IpsetName, "src", "-i", infName, "-j", "ACCEPT")

	// drop all
	_ = RunCmd("iptables", "-P", "INPUT", "DROP")
	return nil
}

func (ipt *IptablesFirewall) Open(srcIp string) error {
	return RunCmd("ipset", "add", IpsetName, srcIp)
}

func (ipt *IptablesFirewall) Close(srcIp string) error {
	return nil
}

// 自动生成代码用的
func _test() {
	var _ Firewall = &IptablesFirewall{}
}
