package mips

import (
//"fmt"
)

import (
	//	"cpu"
	//	"cpu/mips"
	//	"dev/mem"
	//	"dev/char"
	//	"dev/char/ser8250"
	"plat"
)

type PlatMalta struct {
	plat.Plat
}

func maltaInit() error {
	return nil
}

func init() {
	println("Init plat/mips/malta")
	//plat.Register("malta", maltaInit())
}

func (pm *PlatMalta) Init() error {
	return nil
}
func (pm *PlatMalta) Start() {
}
