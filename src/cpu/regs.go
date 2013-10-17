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
	GetVal() uint64
	GetName() string
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

func (r *Gpr) GetVal() uint64        { return r.val }
func (r *Gpr) GetName() string       { return r.name }
func (r *Gpr) SetAccess(t RegAccess) { r.access = t }
func (r *Gpr) SetVal(v uint64)       { r.val = v }
func (r *Gpr) SetName(s string)      { r.name = s }

// UpdateFields is to generate individual fields from Reg.Val
// This is only called when theres no fields,
// Specific register need not implement this
func (r *Gpr) UpdateFields(v uint64) error { return r.SetVal(v) }

// UpdateReg updates Reg.Val from specific fields. This function does
// the opposite of UpdateFields
func (r *Gpr) UpdateReg() (e error) { return }

type SpclReg struct { // Special Register
	Gpr
	resetVal  uint64
	rsrvdOne  uint64 // Read as One
	rsrvdZero uint64 // Read as Zero,
	Valid     bool
}

func (r *SpclReg) SetVal(v uint64) (e error) {
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
	v &= ^r.resrvdZero

	r.val = v

	return
}

type CopReg struct {
	SpclReg
}

type FpReg struct {
	SpclReg
}
