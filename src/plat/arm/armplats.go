package arm

import (
	//"fmt"
	"errors"
)

import (
	"cpu/arm"
	"plat"
	"util"

	//	_ "plat/arm"
	//	_ "plat/arm/other"
)

type PlatArm struct {
	plat.Plat
}

func (p *PlatArm) StartOne(n int) error {
	if n > p.NumCores {
		return errors.New("Index out of range")
	}
	return p.Cores[n].Start()

}
func (p *PlatArm) Start() {
	for i := 0; i < p.NumCores; i++ {
		p.Cores[i].Init()
	}
}

func (p *PlatArm) Init() error {
	util.PrintMe()
	arm.Init()
	return nil
}

func Start() {
	util.PrintMe()
}

func init() {
	util.PrintMe()
	ebInit()
	pbInit()
	//cavium.Init()
}
