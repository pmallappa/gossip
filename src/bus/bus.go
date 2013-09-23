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
	end   uint64
	dr    ReadWriterAll
}

// Bus is implemented as binary tree,
type Bus struct {
	devices []*device
	left    *Bus
	right   *Bus
}

const (
	DEVICE = 1 << iota
	ALIAS
	INVALID
)

const (
	RO = 1 << iota
	WO
	RRO //Raw Read Only
	RWO // Raw Read Write
	RW  = RO | WO
	RRW = RRO | RWO
)

type Reader interface {
	Read8At(off uint64) (uint8, error)
	Read16At(off uint64) (uint16, error)
	Read32At(off uint64) (uint32, error)
	Read64At(off uint64) (uint64, error)
}

type Writer interface {
	Write8At(off uint64, val uint8) error
	Write16At(off uint64, val uint16) error
	Write32At(off uint64, val uint32) error
	Write64At(off uint64, val uint64) error
}

type ReaderAt interface {
	ReadAt(p []byte, off uint64) (n int, e error)
}

type WriterAt interface {
	WriteAt(p []byte, off uint64) (n int, e error)
}

type ReadWriter interface {
	Reader
	Writer
}

type ReadWriterAt interface {
	ReaderAt
	WriterAt
}

type ReadWriterAll interface {
	ReadWriterAt
	ReadWriter
}

func (b *Bus) Read8At(addr uint64) (uint8, error) {
	// Since we are reading a byte, no need to 'endianize'
	var err error
	if dr, off, err := b.getDevice(addr); err == nil {
		return dr.Read8At(off)
	}

	return 0, err
}

func (b *Bus) Read16At(addr uint64) (uint16, error) {
	var err error
	if dr, off, err := b.getDevice(addr); err == nil {
		return dr.Read16At(off)
	}
	return 0, err
}

func (b *Bus) Read32At(addr uint64) (uint32, error) {
	var err error
	if dr, off, err := b.getDevice(addr); err == nil {
		return dr.Read32At(off)
	}
	return 0, err
}

func (b *Bus) Read64At(addr uint64) (uint64, error) {
	var err error
	if dr, off, err := b.getDevice(addr); err == nil {
		return dr.Read64At(off)
	}
	return 0, err
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

func (b *Bus) Write8At(val uint8, addr uint64) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}
	return dr.Write8At(off, val)
}

func (b *Bus) Write16At(addr uint64, val uint16) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}

	return dr.Write16At(off, val)
}

func (b *Bus) Write32At(addr uint64, val uint32) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}
	return dr.Write32At(off, val)
}

func (b *Bus) Write64At(addr uint64, val uint64) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}

	return dr.Write64At(off, val)
}

func (b *Bus) ReadAt(buf []byte, addr uint64) (int, error) {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return 0, err
	}
	return dr.ReadAt(buf, off)
}

func (b *Bus) WriteAt(buf []byte, addr uint64) (int, error) {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return 0, err
	}
	return dr.WriteAt(buf, off)
}

func (b *Bus) AddDevice(addr, size uint64, rw ReadWriterAll) error {
	if _, _, err := b.getDevice(addr); err == nil {
		b.add(addr, size, rw)
	}
	return nil
}
