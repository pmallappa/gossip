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
	ReadAt8(val *uint8, off uint64) error
	ReadAt16(val *uint16, off uint64) error
	ReadAt32(val *uint32, off uint64) error
	ReadAt64(val *uint64, off uint64) error
}

type Writer interface {
	WriteAt8(val uint8, off uint64) error
	WriteAt16(val uint16, off uint64) error
	WriteAt32(val uint32, off uint64) error
	WriteAt64(val uint64, off uint64) error
}

type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, e error)
}

type WriterAt interface {
	WriteAt(p []byte, off int64) (n int, e error)
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

func (b *Bus) ReadAt8(val *uint8, addr uint64) error {
	// Since we are reading a byte, no need to 'endianize'
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}
	return dr.ReadAt8(val, off)
}

func (b *Bus) ReadAt16(val *uint16, addr uint64) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}
	return dr.ReadAt16(val, off)
}

func (b *Bus) ReadAt32(val *uint32, addr uint64) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}
	return dr.ReadAt32(val, off)
}

func (b *Bus) ReadAt64(val *uint64, addr uint64) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}
	return dr.ReadAt64(val, off)
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

func (b *Bus) WriteAt8(val uint8, addr uint64) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}
	return dr.WriteAt8(val, off)
}

func (b *Bus) WriteAt16(val uint16, addr uint64) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}

	return dr.WriteAt16(val, off)
}

func (b *Bus) WriteAt32(val uint32, addr uint64) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}
	return dr.WriteAt32(val, off)
}

func (b *Bus) WriteAt64(val uint64, addr uint64) error {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return err
	}

	return dr.WriteAt64(val, off)
}

func (b *Bus) ReadAt(buf []byte, addr uint64) (int, error) {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return 0, err
	}
	return dr.ReadAt(buf, int64(off))
}

func (b *Bus) WriteAt(buf []byte, addr uint64) (int, error) {
	dr, off, err := b.getDevice(addr)
	if err != nil {
		return 0, err
	}
	return dr.WriteAt(buf, int64(off))
}

func (b *Bus) AddDevice(addr, size uint64, rw ReadWriterAll) error {
	if _, _, err := b.getDevice(addr); err == nil {
		b.add(addr, size, rw)
	}
	return nil
}
