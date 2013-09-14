package mips

import (
//"fmt"
)

// import cpu/mips as mipscpu
// import plat/mips as mipsplat

import (
	//mipscpu //"cpu/mips"
	"plat"
)

type PlatMips struct {
	plat.Plat
}

func (p *PlatMips) Start() {
	for i := 0; i < p.NumCores; i++ {
		go p.Cores[i].Start()
	}
}

func ParseFlags() error {
	var opts map[string]string
	var e error
	if opts, e = plat.ParseFlags(); e != nil {
		return e
	}
	for k, v := range opts {
		println(k, v)
	}
	return nil
}

func (p *PlatMips) Init() error {
	ParseFlags()
	for i := 0; i < p.NumCores; i++ {
		p.Cores[i].Init()
	}
	return nil
}

func (p *PlatMips) Setup() error {
	return nil
}
