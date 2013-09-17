// Package char is implentation of all charatecter devices
// For a given device on a platform, this package provices corresponding host
// implementation, this part takes care of how the data from platform device is
// communicated to the world.
package char

import (
	"io"
	)

import (
	"dev"
)

type CDevHostType uint32

const (
	CDevUDP CDevHostType = 1 << (8 + iota)
	CDevPIPE
	CDevFILE
	CDevVC
	CDevCONSOLE
	CDevSOCK
	CDevPARPORT
	CDevTTY
)

type CharDev struct {
	ctype		CDevHostType
	id			uint16
	mux			bool

	dev.Device

	path		string

	host		string
	port		uint16
	localaddr	uint64
	localport	uint32

}

// Parse all that we can, rest will be passed it down to the actual device module
func (c *CharDev) ParseFlags(m map[string]string) (map[string]string, error) {
	return m, nil
}