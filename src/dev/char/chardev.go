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

type CdevHost uint32

const (
	CHOST_UDP CdevHost = 1 << (8 + iota)
	CHOST_PIPE
	CHOST_FILE
	CHOST_VC
	CHOST_CONSOLE
	CHOST_SOCK
	CHOST_PARPORT
	CHOST_TTY
)

type CharDev struct {
	ctype		CdevHost
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
func ParseFlags(m map[string]string) (map[string]string, error) {
	return m, nil
}