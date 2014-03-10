package cpu

import (
	"util"
)

var availableCpu []*Cpu

func init() {
	availableCpu = make([]*Cpu, 0, 64)
	util.PrintMe()
	initCpuOpts()
}

func Init() error {

	return nil
}

func RegisterCpu(cpu *Cpu) error {
	//availableCpu = append(availableCpu, cpu)
	return nil
}
