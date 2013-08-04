// Package Bus
// Implementing a simple bus, doesn't do multiplexing, all read/write operations
// are expected to take once cycle.

package bus

// System
import (
	"encoding/binary"
	//"fmt"
)

import (
	"dev"
)

type Bus struct {
	devices []dev.Device
	endian  binary.ByteOrder
}

const (
	DEVICE = 1 << iota
	ALIAS
	INVALID
)

type BusReader interface {
	Read8(uint32) (uint8, error)
	Read16(uint32) (uint16, error)
	Read32(uint32) (uint32, error)
	Read64(uint32) (uint64, error)
}

type BusRawReader interface {
	RawRead(uint32, size uint32) []byte
}

type BusWriter interface {
	Write8(uint32, b uint8) error
	Write16(uint32, v uint16) error
	Write32(uint32, v uint32) error
	Write64(uint32, v uint64) error
}

type BusRawWriter interface {
	RawWrite(uint32, []byte) error
}

func (b *Bus) SetEndian(o binary.ByteOrder) {
	b.endian = o
}

func (b *Bus) getDevice(addr uint64) (*dev.Device, off) {
	// Need to implement the B-tree to have devices at addresses specified by Map from platform
	return nil, 0
}

func (b *Bus) Read(addr uint64, size uint32) (interface{}, error) {
	var val interface{}
	var dr Reader
	dr, off = getDevice(addr)
	if d != nil {
		switch size >> 3 {
		case 8:
			val, error := dr.Read8(off)
		case 16:
			val, error := dr.Read16(off)
		case 32:
			val, error := dr.Read32(off)
		case 64:
			fallthrough
		default:
			val, error = dr.Read64(off)
		}

		if error != nil {
		}
	}
}

func (b *Bus) Read8(addr uint64) (uint8, error) {
	// We have a situation; accessing the wrong address on the bus
	//
	dr, off, err := getDevice(addr)
	if err != nil {
		return 0, err
	}

	// Since we are reading a byte, no need to 'endianize'
	return dr.Read8(off)
}

func (b *Bus) Read16(addr uint64) (uint16, error) {
	var val uint16 = 0
	dr, off, err := getDevice(addr)
	if err != nil {
		return 0, err
	}
	val, err := dr.Read16(off)
	return b.endian.Uint16(val), err
}

func (b *Bus) Read32(addr uint64) (uint32, error) {
	var val uint32 = 0
	dr, off, err := getDevice(addr)
	if err != nil {
		return 0, err
	}
	val, err := dr.Read32(off)

	return b.endian.Uint32(val), err
}

func (b *Bus) Read64(addr uint64) (uint64, error) {
	var val uint64 = 0
	dr, off, err := getDevice(addr)
	if err != nil {
		return 0, err
	}
	val, err := dr.Read64(off)

	return b.endian.Uint64(val), err
}

func (b *Bus) Write8(addr uint64, val uint8) error {
	dr, off, err := getDevice(addr)
	if err != nil {
		return err
	}
	return dr.Write8(off, val)
}

func (b *Bus) Write16(addr uint64, v uint16) error {
	var wval uint16
	dr, off, err := getDevice(addr)
	if err != nil {
		return err
	}

	m.endian.PutUint16(m.buf[addr:], wval)
	return dr.Write16()
}

func (b *Bus) Write32(addr uint64, v uint32) error {
	m.endian.PutUint32(m.buf[addr:], v)
	return nil
}

func (b *Bus) Write64(addr uint64, v uint64) error {
	m.endian.PutUint64(m.buf[addr:], v)
	return nil
}

func (b *Bus) RawRead(addr uint64, size uint32) []byte {
	return m.buf[addr : addr+size]
}

func (b *Bus) RawWrite(addr uint64, buf []byte, size uint32) error {
	copy(m.buf[addr:], buf)
	return nil
}
