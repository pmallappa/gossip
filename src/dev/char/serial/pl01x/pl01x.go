package pl01x

import (
	"io"
)

import (
	"dev"
	"dev/char"
	"dev/char/serial"
)

const (
	UARTDR    = 0x00 // Data Register
	UARTRSR   = 0x04 // Receive Status Register
	UARTECR   = 0x04 // Error Clear Register
	UARTFR    = 0x18 // Flag Register
	UARTILPR  = 0x20 // IrDA Low-Power Counter Register
	UARTIBRD  = 0x24 // Integer Baud Rate Register
	UARTFBRD  = 0x28 // Fractional Baud Rate Register
	UARTLCR_H = 0x2C // Line Control Register
	UARTCR    = 0x30 // Control Register
	UARTIFLS  = 0x34 // Interrupt FIFO Level Select Register
	UARTIMSC  = 0x38 // Interrupt Mask Set/Clear Register
	UARTRIS   = 0x3C // Interrupt Status Register
	UARTMIS   = 0x40 // Interrupt Status Register
	UARTICR   = 0x44 // Clear Register
	UARTDMACR = 0x48 // DMA Control Register

)

type pl011 struct {
	regs [0x50]dev.Register
	wr   io.ReadWrite
}

func (p *pl011) Init() error {
	if p.wr == nil {
		return errors.New("No Transmit method defined")
	}

	for i := 0; i < len(p.regs); i++ {
		switch i {
		case UARTDR:
			p.regs[i] = &regDR
			p.regs[i].wr = p.wr
		case UARTRSR, UARTECR,
			UARTFR, UARTILPR, UARTIBRD,
			UARTFBRD, UARTLCR_H, UARTCR,
			UARTIFLS, UARTIMSC, UARTRIS,
			UARTMIS, UARTICR, UARTDMACR:
			p.regs[i] = dev.NewReg()
		default:
			p.regs[i] = ZeroReg
		}

	}

}

type regDR struct {
	fifolen u8
	fifo    [16]u8
	wr      io.ReadWriter
}

func (pl *pl011) Xmit() {

}

func (pl *pl011) write32(off uint32, val uint32) error {
	switch off {
	case UARTDR:
		p.regDR.write(val)
	case UARTRSR:
		fallthrough
	case UARTECR:
	case UARTFR:
	case UARTILPR:
	case UARTIBRD:
	case UARTFBRD:
	case UARTLCR_H:
	case UARTCR:
	case UARTIFLS:
	case UARTIMSC:
	case UARTRIS:
	case UARTMIS:
	case UARTICR:
	case UARTDMACR:
	default:

	}
}

func (pl *pl011) Write(off uint32, val uint32) error {
	return write32(off, val)
}

func (pl *pl011) read32(off uint32) (uint32, error) {

}

func (pl *pl011) read8(off uint32) (uint8, error) {
	val, err := read32(off)

	return uint8(val), err
}

func (pl *pl011) read16(off uint32) (uint16, error) {
	val, err := read32(off)

	return uint16(val), err
}

func (pl *pl011) Read(off uint32) (uint32 error) {
	return read32(off)
}
