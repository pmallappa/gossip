// Package Mem,
// Not much to do read/write bytes

package mem

// System
import (
	"encoding/binary"
	//"fmt"
)

// local
import (
//"util"
)

type Mem struct {
	buf    []byte
	endian binary.ByteOrder
}

func (m *Mem) SetEndian(o binary.ByteOrder) {
	m.endian = o
}

func (m *Mem) Read(off uint32) (uint8, error) {
	return m.buf[off], nil
}

func (m *Mem) Read16(off uint32) (uint16, error) {
	//var v uint32
	return m.endian.Uint16(m.buf[off:]), nil
}

func (m *Mem) Read32(off uint32) (uint32, error) {
	return m.endian.Uint32(m.buf[off:]), nil
}

func (m *Mem) Read64(off uint32) (uint64, error) {
	return m.endian.Uint64(m.buf[off:]), nil
}

func (m *Mem) Write(off uint32, b uint8) error {
	m.buf[off] = b
	return nil
}

func (m *Mem) Write16(off uint32, v uint16) error {
	m.endian.PutUint16(m.buf[off:], v)
	return nil
}

func (m *Mem) Write32(off uint32, v uint32) error {
	m.endian.PutUint32(m.buf[off:], v)
	return nil
}

func (m *Mem) Write64(off uint32, v uint64) error {
	m.endian.PutUint64(m.buf[off:], v)
	return nil
}

func (m *Mem) RawRead(off uint32, size uint32) []byte {
	return m.buf[off : off+size]
}

func (m *Mem) RawWrite(off uint32, buf []byte, size uint32) error {
	copy(m.buf[off:], buf)
	return nil
}

func Newmem(size uint64) *Mem {
	m := new(Mem)
	m.buf = make([]byte, size)
	return m
}
