package cpu

import (
	"flag"
	"log"
)

import (
	"util"
)

var cpu_opts string

var availableCpu []*CpuInfo

func init() {
	availableCpu = make([]*CpuInfo, 0, 64)
	util.PrintMe()
	flag.StringVar(&cpu_opts, "cpu", "",
		"CPU's, type ? to list, but -plat should be provided")
}

func InitGeneric() {
	util.PrintMe()
}

func RegisterCpu(cpu *CpuInfo) {
	availableCpu = append(availableCpu, cpu)
}

func SetLogger(l *log.Logger) {
	l = l
}
