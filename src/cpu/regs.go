package cpu

import (
	"fmt"
)

import (
//	"util"
)

type RegAccess uint8

// By Default All Registers are Read/Write
const (
	RDWR RegAccess = 0
	INVALID
	RDONLY = 1 << iota
	WRONLY
	RESERVED
)

type Register interface {
	SetVal(uint64) error
	GetVal() (uint64, error)
	GetName() string
	SetName(string) error
	UpdateFields(uint64) error
	UpdateReg() error
}

// Co processor Registers
type Reg struct {
	Name         string
	Val          uint64
	ResetVal     uint64
	ReservedOne  uint64    // Read as One
	ReservedZero uint64    // Read as Zero,
	access       RegAccess // RDRW,RDONLY,RDRW
}

func (r *Reg) GetVal() (uint64, error) {
	val := r.Val
	val |= r.ReservedOne
	val &= ^r.ReservedZero
	return val, nil
}
func (r *Reg) GetName() string {
	return r.Name
}
func (r *Reg) SetName(s string) error {
	r.Name = s
	return nil
}

func (r *Reg) SetVal(v uint64) error {
	r.Val = v
	return nil
}

// UpdateFields is to generate individual fields from Red.Val
// This is only called when theres no fields,
// Specific register need not implement this
func (r *Reg) UpdateFields(uint64) error {
	return nil
}

// UpdateReg updates Reg.Val from specific fields. This function does
// the opposite of UpdateFields
func (r *Reg) UpdateReg() error {
	return nil
}
func (r *Reg) SetAccess(t RegAccess) {
	r.access = t
}

func (r *CopReg) SetVal(v uint64) error {
	var e error = nil
	if r.ReservedZero&v != 0 {
		e = fmt.Errorf("writing to reserved fields of %s field(s): %x",
			r.Name, r.ReservedZero&v)
	}
	if r.ReservedOne&v != 0 {
		e = fmt.Errorf("writing to reserved fields of %s field(s): %x",
			r.Name, r.ReservedOne&v)
	}
	// We will simulate hardware hard-wired bits.
	// RAZ and RAO will not be changed
	v |= r.ReservedOne
	v &= ^r.ReservedZero

	r.Val = v

	return e
}

type CoreReg struct {
	Reg
	Alias string
}

type CopReg struct {
	Reg
	Valid bool
}

type FpReg struct {
	Reg
}
