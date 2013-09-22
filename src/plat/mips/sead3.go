package mips

import (
//"fmt"
)

import (
	//"cpu/mips"
	//"dev/mem"
	//"dev/char"
	//"dev/char/ser8250"
	"plat"
)

type PlatSead3 struct {
	PlatMips
}

func sead3Init() error {
	return nil
}

func init() {
	//println("Init plat/mips/sead3")
	plat.Register(NewPlatSead3())
}

// PlatController
func (pm *PlatSead3) Init() error {
	return nil
}

func (pm *PlatSead3) Start() error {
	return nil
}

func (pm *PlatSead3) Stop() error {
	return nil
}

func (pm *PlatSead3) ParseFlags() error {
	return nil
}

// PlatCustomizer interface
func (pm *PlatSead3) PreSetup() error {
	return nil
}

func (pm *PlatSead3) CustomSetup() error {
	return nil
}

func (pm *PlatSead3) PostSetup() error {
	return nil
}

func NewPlatSead3() *PlatSead3 {
	p := new(PlatSead3)
	p.SetInfo("sead3", "mips", "1.0")
	return p
}

// PlatDebugger
func (pm *PlatSead3) Pause() error {
	return nil
}

func (pm *PlatSead3) Resume() error {
	return nil
}
