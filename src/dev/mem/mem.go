// Package Mem,
// read/write bytes, halfwords, workds, and boublewords

package mem

// System
import (
	"encoding/binary"
	"errors"
	//"fmt"
)

// local
import (
//"util"
)

// Each CPU/Core gets its own read/write channel
// Address are communicated through 'addr' channel.
// Response are sent via the read/write channel.
type Mem struct {
	buf    []byte
	endian binary.ByteOrder
	rd     []chan uint64
	wr     []chan uint64
	addr   []chan uint64
}

func (m *Mem) SetEndian(o binary.ByteOrder) {
	m.endian = o
}

func (m *Mem) Read8At(off uint64) (uint8, error) {
	return m.buf[off], nil
}

func (m *Mem) Read16At(off uint64) (uint16, error) {
	return m.endian.Uint16(m.buf[off:]), nil
}

func (m *Mem) Read32At(off uint64) (uint32, error) {
	return m.endian.Uint32(m.buf[off:]), nil
}

func (m *Mem) Read64At(off uint64) (uint64, error) {
	return m.endian.Uint64(m.buf[off:]), nil
}

func (m *Mem) Write8At(off uint64, val uint8) error {
	m.buf[off] = val
	return nil
}

func (m *Mem) Write16At(off uint64, val uint16) error {
	m.endian.PutUint16(m.buf[off:], val)
	return nil
}

func (m *Mem) Write32At(off uint64, val uint32) error {
	m.endian.PutUint32(m.buf[off:], val)
	return nil
}

func (m *Mem) Write64At(off uint64, val uint64) error {
	m.endian.PutUint64(m.buf[off:], val)
	return nil
}

func (m *Mem) ReadAt(p []byte, off uint64) (n int, e error) {
	n = copy(p, m.buf[off:])
	if n != len(p) {
		e = errors.New("Read: Access Outside Memory")
	}
	return
}

// The bus will call this function, so does CPU,
// if
func (m *Mem) WriteAt(p []byte, off uint64) (n int, e error) {
	n = copy(m.buf[off:], p)
	if n != len(p) {
		e = errors.New("Read: Access Outside Memory")
	}
	return
}

func Newmem(size uint64) *Mem {
	m := new(Mem)
	m.buf = make([]byte, size)
	return m
}

func Init() {

}
