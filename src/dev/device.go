package dev

// system imports
import (
//"encoding/binary"
//"fmt"
)

// local imports
import (
//"util"
)

// Each device has a host part and device part
// Device part tells how to present itself to an operating system
// running on simulator, Host part tells how to transfer the
// data in-and-out of the simulated machine

type DEVtype uint32

const (
	CHAR DEVtype = 1 << iota
	NET
	BLK
	PCI
	USB
)

// All devices recieve the offset to read the memory controller strips-off
// the base and only forwards offset to read.
// Any Device can interrupt,
// Interrupt channel is sent via Init() or Requesting a new device.
type InterrupterEdge interface {
	AssertEdge(int) error
	DeassertEdge(int) error
}

type InterrupterLevel interface {
	DeassertLevel(int) error
	AssertLevel(int) error
}

type DevInfo struct {
	model   string
	vendor  string
	version string
}

func (d *DevInfo) GetInfo() map[string]string {
	return map[string]string{"model": d.model, "vendor": d.vendor, "version": d.version}
}

type Dev struct {
	devtype DEVtype
	//regs     []byte
	irq      uint
	intrctrl chan bool
}

type Device struct {
	DevInfo
	Dev
}

// All devices must implement bus.Reader bus.Writer bus.RawReader bus.RawWriter

type Devicer interface {
	Initialize() error
}

type Reader interface {
	Read(uint64) (uint64, error)
	Read8(uint64) (uint8, error)
	Read16(uint64) (uint16, error)
	Read32(uint64) (uint32, error)
}

type Writer interface {
	Write(uint64, uint64) error
	Write8(uint64, uint8) error
	Write16(uint64, uint16) error
	Write32(uint64, uint32) error
}

type ReadWriter interface {
	Reader
	Writer
}

func NewDevice(size uint64) *Device {
	m := new(Device)
	//m.regs = make([]byte, size)
	return m
}
