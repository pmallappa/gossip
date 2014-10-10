package serial

import (
	//"dev"
	"dev/char"
)

type Serial interface {
	char.CharDevice
}

type Ser struct {
	char.CharDev
}
