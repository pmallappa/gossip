package cpu

import (
	"flag"
)

import (
	"util"
)

var cpu_opts string

var availableCpu []*Core

func init() {
	availableCpu = make([]*Core, 0, 64)
	util.PrintMe()
	flag.StringVar(&cpu_opts, "cpu", "",
		"CPU's, type ? to list, but -plat should be provided")
}

func InitGeneric() {
	util.PrintMe()
}

func RegisterCpu(cpu *Core) {
	//availableCpu = append(availableCpu, cpu)
}
