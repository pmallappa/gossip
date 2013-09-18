// Package Bus
// Implementing a simple bus, doesn't do multiplexing, all read/write operations
// are expected to take once cycle.

package bus

// System
import (
//"fmt"
)

import (
//"dev"
)

type device struct {
	start uint64
	size  uint64
	endi	uint64
	dr    ReadWriterAll
}

// Bus is implemented as binary tree,
type Bus struct {
	dev *device
	left *Bus
	right *Bus
}

const (
	DEVICE = 1 << iota
	ALIAS
	INVALID
)

const (
	RO = 1 << iota
	WO
	RRO	//Raw Read Only
	RWO	// Raw Read Write
	RW  = RO | WO
	RRW = RRO | RWO
)

type Reader interface {
	Read8(uint64) (uint8, error)
	Read16(uint64) (uint16, error)
	Read32(uint64) (uint32, error)
	Read64(uint64) (uint64, error)
}

type Writer interface {
	Write8(uint64, uint8) error
	Write16(uint64, uint16) error
	Write32(uint64, uint32) error
	Write64(uint64, uint64) error
}

type RawReader interface {
	RawRead(uint64, []byte) error
}

type RawWriter interface {
	RawWrite(uint64, []byte) error
}

type ReadWriter interface {
	Reader
	Writer
}

type RawReadWriter interface {
	RawReader
	RawWriter
}

type ReadWriterAll interface {
	ReadWriter
	RawReadWriter
}

func (b *Bus) getDevice(addr uint64) (ReadWriterAll, uint64, error) {
	// Need to implement the B-tree to have devices at
	// addresses specified by Map from platform
	
	return b.dev[0].dr, 0, nil
}

func (b *Bus) Read8(addr uint64) (uint8, error) {
	// Since we are reading a byte, no need to 'endianize'
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return 0, err
	}
	return dr.Read8(off)
}

func (b *Bus) Read16(addr uint64) (uint16, error) {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return 0, err
	}
	return dr.Read16(off)
}

func (b *Bus) Read32(addr uint64) (uint32, error) {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return 0, err
	}
	return dr.Read32(off)
}

func (b *Bus) Read64(addr uint64) (uint64, error) {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return 0, err
	}
	return dr.Read64(off)
}

/*
 * Can have generic function like following.
 * This is going to cause additional runtime check, which may turn out to be
 * expensive
 */
/*

func (b *Bus) ReadNG(addr uint64, size uint32) (interface{}, error) {
	var val interface{}
	var err error
	dr, off = getDevice(addr)
	if d != nil {
		switch size >> 3 {
		case 8:
			val, err = dr.Read8(off)
		case 16:
			val, err = dr.Read16(off)
		case 32:
			val, err = dr.Read32(off)
		case 64:
			fallthrough
		default:
			val, err = dr.Read64(off)
		}

		if error != nil {
		}
	}
}

func (b *Bus) WriteNG(addr uint64, val interface{}) error {
	dr, off, err := getDevice(addr)
	if err != nil {
		return err
	}
	switch val.(type) {
	case uint8:
		dr.Write8(off, val)
		return dr.Write8(off, val)
	case uint16:
		dr.Write16(off, val)
		return dr.Write16(off, val)
	case uint32:
		dr.Write32(off, val)
		return dr.Write32(off, val)
	case uint64:
		dr.Write64(off, val)
		return dr.Write64(off, val)

	}

}
*/

func (b *Bus) Write8(addr uint64, val uint8) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}
	return dr.Write8(off, val)
}

func (b *Bus) Write16(addr uint64, val uint16) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}

	return dr.Write16(off, val)
}

func (b *Bus) Write32(addr uint64, val uint32) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}
	return dr.Write32(off, val)
}

func (b *Bus) Write64(addr uint64, val uint64) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}

	return dr.Write64(off, val)
}

func (b *Bus) RawRead(addr uint64, buf []byte) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}
	return dr.RawRead(off, buf)
}

func (b *Bus) RawWrite(addr uint64, buf []byte) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}
	return dr.RawWrite(off, buf)
}



func (b *Bus)AddDevice(addr, size uint64, rw ReadWriterAll) error {
	if dr, off, err := b.getDevice(addr); err == nil {
		b.add(&device{addr, size, rw})
	}

	return err
}

