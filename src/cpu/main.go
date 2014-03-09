package cpu

import (
	"util"
	"util/cflag"
)

var cpu_opts = cflag.NewCFlagSet("cpu")

var availableCpu []*Cpu

func init() {
	availableCpu = make([]*Cpu, 0, 64)
	util.PrintMe()
	//flag.StringVar(&cpu_opts, "cpu", "",
	//	"CPU's, type ? to list, but -plat should be provided")
}

func Init() error {

	return nil
}

func RegisterCpu(cpu *Cpu) error {
	//availableCpu = append(availableCpu, cpu)
	return nil
}
