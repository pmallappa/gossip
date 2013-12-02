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
	RDWR RegAccess = iota
	RDONLY
	WRONLY
	RESERVED
	INVALID
)

type Register interface {
	Set(uint64) error
	Val() uint64
	Name() string
	SetName(string)
	UpdateFields(uint64) error
	UpdateReg() error
}

// Co processor Registers
type Gpr struct {
	name   string
	val    uint64
	access RegAccess // RDRW,RDONLY,RDRW
}

func (r *Gpr) Val() uint64                                 { return r.val }
func (r *Gpr) Name() string                                { return r.name }
func (r *Gpr) SetAccess(t RegAccess)                       { r.access = t }
func (r *Gpr) Set(v uint64) error                          { r.val = v; return nil }
func (r *Gpr) SetName(s string)                            { r.name = s }
func (r *Gpr) SetAll(name string, val uint64, a RegAccess) { r.name = name; r.val = val; r.access = a }

// UpdateFields is to generate individual fields from Reg.Val
// This is only called when theres no fields,
// Specific register need not implement this
func (r *Gpr) UpdateFields(v uint64) error { r.Set(v); return nil }

// UpdateReg updates Reg.Val from specific fields. This function does
// the opposite of UpdateFields
func (r *Gpr) UpdateReg() (e error) { return }

type SpclReg struct { // Special Register
	Gpr
	resetVal  uint64
	rsrvdOne  uint64 // Read as One
	rsrvdZero uint64 // Read as Zero,
	valid     bool
}

func (r *SpclReg) SetReserved(mask uint64, ones bool) {
	if ones {
		r.rsrvdOne = mask
	} else {
		r.rsrvdZero = mask
	}
}

func (r *SpclReg) Valid() bool { return r.valid }
func (r *SpclReg) Set(v uint64) (e error) {
	if r.rsrvdZero&v != 0 {
		e = fmt.Errorf("writing to reserved fields of %s field(s): %x",
			r.Name, r.rsrvdZero&v)
	}
	if r.rsrvdOne&v != 0 {
		e = fmt.Errorf("writing to reserved fields of %s field(s): %x",
			r.Name, r.rsrvdOne&v)
	}
	// We will simulate hardware hard-wired bits.
	// RAZ and RAO will not be changed
	v |= r.rsrvdOne
	v &= ^r.rsrvdZero

	r.val = v

	return
}

func NewSpclReg(resetval uint64, valid bool, name string) *SpclReg {
	return &SpclReg{
		Gpr:      Gpr{name: name},
		resetVal: resetval,
		valid:    true,
	}
}

type CopReg struct {
	SpclReg
}

type FpReg struct {
	SpclReg
}
