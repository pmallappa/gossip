package dev

// system imports
import (
//"encoding/binary"
//"flag"

//"fmt"
)

// local imports
import (
	"bus"
	"util/cflag"
)

// Each device has a host part and device part
// Device part tells how to present itself to an operating system
// running on simulator, Host part tells how to transfer the
// data in-and-out of the simulated machine

type Type uint32

const (
	CHAR Type = 1 << iota
	NET
	BLK
	MISCDEV
)

type Bustype uint32

const (
	PCI Bustype = 1 << iota
	I2C
	SCSI
	IDE
	HDA
	USB
	MISCDEVBUS
)

// All devices recieve the offset to read the memory controller strips-off
// the base and only forwards offset to read.
// Any Device can interrupt,
// Interrupt channel is sent via Init() or Requesting a new device.
type EdgeInterrupt interface {
	AssertEdge(int) error
	DeassertEdge(int) error
}

type LevelInterrupt interface {
	AssertLevel(int) error
	DeassertLevel(int) error
}

type Interrupter interface {
	SetInterrupt(chan bool)
}

type _Info struct {
	model  string
	vendor string
	id     string
}

func (i *_Info) GetInfo() map[string]string {
	return map[string]string{"model": i.model, "vendor": i.vendor, "id": i.id}
}

type _Dev struct {
	devtype   DEVtype
	irq       uint
	interrupt chan bool
	opts      cflag.OptionT // Options specific to this device
}

type Dev struct {
	_Info
	_Dev
}

// All Devices must support all Read/Write
type Device interface {
	bus.ReadWriterAll
	Interrupter
	Initializer
}

// All devices must implement bus.ReadWriterAll
type Initializer interface {
	Init() error
	Configure() error
}

type Parser interface {
	Parse(string) error
}

func NewDevice(size uint64) *Dev {
	m := new(Dev)
	//m.regs = make([]byte, size)
	return m
}

func RegisterDev() {}

func init() {
	initDevFlags()
}

type RegAccess uint8

const (
	R_RD RegAccess = iota
	R_WR
	R_RDWR = R_RD | R_WR
	R_RESERVED
	R_INVALID
)

type DevReg struct {
	name   string
	val    uint32    // could be uint8/16/32
	access RegAccess // RD,RDWR,WR
}
