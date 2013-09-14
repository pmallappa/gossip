package mips

import (
//"fmt"
)

import (
	//"cpu/mips"
	//	"dev/mem"
	//	"dev/char"
	//	"dev/char/ser8250"
	"plat"
)

type PlatMalta PlatMips

func maltaInit() error {
	return nil
}

func init() {
	println("Init plat/mips/malta")
	plat.Register(NewPlatMalta())
}

// PlatController
func (pm *PlatMalta) Init() error {
	return nil
}

func (pm *PlatMalta) Start() error {
	return nil
}

func (pm *PlatMalta) Stop() error {
	return nil
}

func (pm *PlatMalta) ParseFlags() error {
	return nil
}

// PlatCustomizer interface
func (pm *PlatMalta) PreSetup() error {
	return nil
}

func (pm *PlatMalta) CustomSetup() error {
	return nil
}

func (pm *PlatMalta) PostSetup() error {
	return nil
}

func NewPlatMalta() *PlatMalta {
	p := new(PlatMalta)
	p.SetInfo("malta", "mips", "1.0")
	return p
}

// PlatDebugger
func (pm *PlatMalta) Pause() error {
	return nil
}

func (pm *PlatMalta) Resume() error {
	return nil
}
