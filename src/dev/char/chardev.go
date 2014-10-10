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

type Type uint32

const (
	CDevT_UDP Type = 1 << (8 + iota)
	CDevT_PIPE
	CDevT_FILE
	CDevT_VC
	CDevT_CONSOLE
	CDevT_SOCK
	CDevT_PARPORT
	CDevT_TTY
)

type CharDevice interface {
	dev.Device
}

type CharDev struct {
	ctyp Type
	id   uint16
	mux  bool

	dev.Dev
	w io.ReadWriter
}

// Parse all that we can, rest will be passed it down to the actual device module
func (c *CharDev) ParseFlags(m map[string]string) (map[string]string, error) {
	return m, nil
}
