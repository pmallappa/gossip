package mips

import (
	"debug/elf"

	//"fmt"
)

// import cpu/mips as mipscpu
// import plat/mips as mipsplat

import (
	//"cpu"
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

func (p *PlatMips) Init() error {
	for i := 0; i < p.NumCores; i++ {
		p.Cores[i].Init()
	}
	return nil
}

func (p *PlatMips) Setup() error {
	return nil
}

// PlatELFLoader interface

func (p *PlatMips) ELFClass() []elf.Class {
	c := make([]elf.Class, 0, 16)
	c = []elf.Class{
		elf.ELFCLASS32,
		elf.ELFCLASS64,
	}
	c = append(c, elf.ELFCLASS32)
	c = append(c, elf.ELFCLASS64)
	return c
}

func (p *PlatMips) ELFMachine() []elf.Machine {
	m := make([]elf.Machine, 0, 16)
	m = append(m, elf.EM_MIPS)
	m = append(m, elf.EM_MIPS_RS3_LE)
	m = append(m, elf.EM_MIPS_RS4_BE)
	return m
}
