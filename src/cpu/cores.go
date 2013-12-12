package cpu

import (
	//"fmt"
	"strconv"
)

import (
	"util/logng"
)

type ExcptType uint32

// Types of Cpu Exceptions
type Exception interface {
	Error() string
	String() string
	Type() ExcptType
}

// 	Type  ExcptType
// 	instr string
// }

type InfoT struct {
	freq   uint32 // Hz
	vendor string // Eg, TI, Qualcomm, NetLogic
	model  string // Eg, OMAP3, SnapDragon, XLP
}

type CoreT struct {
	logng.LogNG
	id    uint32 // SMP ID
	cycle uint64 // Processor cycle, modified after every instruction
	excpt Exception
}

type Core struct {
	InfoT
	CoreT
}

func (c *InfoT) Freq() uint64        { return uint64(c.freq) }
func (c *InfoT) SetFreq(freq uint64) { c.freq = uint32(freq) }
func (c *InfoT) Info() map[string]string {
	return map[string]string{
		"model":  c.model,
		"vendor": c.vendor,
		"freq":   strconv.FormatUint(uint64(c.freq), 10),
	}
}

func (c *InfoT) SetInfo(vendor string, model string) {
	c.vendor = vendor
	c.model = model

}

func (c *CoreT) ID() uint32    { return c.id } // Return CPU ID
func (c *CoreT) Cycle() uint64 { return c.cycle }
func (c *CoreT) Setup() error  { return nil }

type CpuController interface {
	Init() error
	Start() error
	Stop() error
}

type CpuStats interface {
	PrintStats() (string, error)
	PrintRegs() (string, error)
}

type CpuExectuter interface {
	Fetch() error
	Decode() error
	Execute() error
}

type CpuInterrupter interface {
	InterruptRaise(uint32) error
	InterruptAck(uint32) error
}

type Cpu interface {
	CpuInterrupter
	CpuExectuter
	CpuStats
	CpuController
	Info() map[string]string
}

type InstrType uint32

// These are internal e(si)mulator errors
type CpuError struct {
	Op  string
	Err error
}
