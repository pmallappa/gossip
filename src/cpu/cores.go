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
type Exception struct {
	Type  ExcptType
	instr string
}

type CpuInfo struct {
	freq   uint32 // Hz
	vendor string // Eg, TI, Qualcomm, NetLogic
	model  string // Eg, OMAP3, SnapDragon, XLP
}

type CpuCore struct {
	logger logng.LoggerNG
	id     uint32 // SMP ID
	cycle  uint64 // Processor cycle, modified after every instruction
	excpt  Exception
}

type Cpu struct {
	CpuInfo
	CpuCore
}

func (c *CpuInfo) SetFreq(freq uint64) {
	c.freq = uint32(freq)
}

func (c *CpuInfo) GetInfo() map[string]string {
	return map[string]string{
		"model":  c.model,
		"vendor": c.vendor,
		"freq":   strconv.FormatUint(uint64(c.freq), 10),
	}
}

func (c *CpuInfo) GetFreq() uint64 {
	return uint64(c.freq)
}
func (c *CpuInfo) SetInfo(vendor string, model string) {
	c.vendor = vendor
	c.model = model
}
func (c *CpuCore) SetLogger(l logng.LoggerNG) {
	c.logger = l
	//c.logger.SetPrefix("CPU" + string(c.id))
}
func (c *CpuCore) GetID() uint32 { // Return CPU ID
	return c.id
}
func (c *CpuCore) GetCycle() uint64 {
	return c.cycle
}

func (c *CpuCore) _getCycles() string {
	return strconv.Itoa(int(c.cycle))
}

func (c *CpuCore) Setup() error {
	//c.logger.SetFn(_getCycles)
	return nil
}

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

type Cores interface {
	CpuInterrupter
	CpuExectuter
	CpuStats
	CpuController
	GetInfo() map[string]string
}

type InstrType uint32

// These are internal e(si)mulator errors
type CpuError struct {
	Op  string
	Err error
}

// Logger interface
func (c *CpuCore) LogLevel(lvl logng.LogLevel, format string,
	v ...interface{}) {
	c.logger.LogLevel(lvl, format, v...)
}

func (c *CpuCore) Log(format string, v ...interface{}) {
	c.logger.Log(format, v...)
}
