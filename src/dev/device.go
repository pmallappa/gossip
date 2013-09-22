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
type EdgeInterrupt interface {
	AssertEdge(int) error
	DeassertEdge(int) error
}

type LevelInterrupt interface {
	DeassertLevel(int) error
	AssertLevel(int) error
}

type Info struct {
	model  string
	vendor string
	id     string
}

func (i *Info) GetInfo() map[string]string {
	return map[string]string{"model": i.model, "vendor": i.vendor, "id": i.id}
}

type Dev struct {
	devtype   DEVtype
	irq       uint
	interrupt chan bool
	rw        bus.ReadWriterAll
	// The options that we couldn't parse
	// may be of some use to the actual device.
	Opts map[string]string
}

type Device struct {
	Info
	Dev
}

// All devices must implement bus.ReadWriterAll
type Devicer interface {
	Initialize() error
	ParseFlags(map[string]string) (map[string]string, error)
}

func NewDevice(size uint64) *Device {
	m := new(Device)
	//m.regs = make([]byte, size)
	return m
}

func init() {
	initDevFlags()
}
