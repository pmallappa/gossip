package cpu

import (
	//"fmt"
	"log"
)

import (
//"util"
)

type ExcptType uint32

// Types of Cpu Exceptions
type Exception struct {
	Type  ExceptType
	instr string
}

// Illegal Instruction and Reserved Instruction means same
const (
	EXCP_RESET ExcptType = iota
	EXCP_RSVDINSTR
	EXCP_MEMACCES // Usually ends up in Pagefault
	EXCP_TLB      // in case of Unified TLB
	EXCP_ITLB
	EXCP_DTLB
	EXCP_DIVZERO // Divide by Zero
	EXCP_BKPT    // Breakpoint
	EXCP_SOFTINT // Way to call Syscall

	EXCP_EXTINT // External Interrupt
	EXCP_MAX    // Processor specific numbering starts from ExceptMax
)

type CpuInfo struct {
	freq   uint32 // Hz
	vendor string // Eg, TI, Qualcomm, NetLogic
	model  string // Eg, OMAP3, SnapDragon, XLP
}

type CpuCore struct {
	logger *log.Logger
	id     uint32 // SMP ID
	cycle  uint64 // Processor cycle, modified after every instruction
	excpt  Exception
	instr  interface{}
}

type Core struct {
	CpuInfo
	CpuCore
}

func (c *CpuInfo) SetFreq(freq uint64) {
	c.freq = uint32(freq)
}
func (c *CpuInfo) GetFreq() uint64 {
	return uint64(c.freq)
}
func (c *CpuInfo) SetInfo(vendor string, model string) {
	c.vendor = vendor
	c.model = model
}
func (c *CpuCore) SetLogger(l *log.Logger) {
	c.logger = l
	c.logger.SetPrefix("CPU" + string(c.id))
}
func (c *CpuCore) GetID() uint32 { // Return CPU ID
	return c.id
}
func (c *CpuCore) GetCycle() uint64 {
	return c.cycle
}

type Cores interface {
	Init() error
	Start() error
	Stop() error
	PrintStats() (string, error)
	PrintRegs() (string, error)
}

type CpuInterrupter interface {
	InterruptRaise(uint32) error
	InterruptAck(uint32) error
}

type InstrType uint32

// These are internal e(si)mulator errors
type CpuError struct {
	Op  string
	Err error
}
