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

type PL011 struct {
	regs    [0x50]dev.Register
	wr      io.ReadWrite
	fifolen u8 // fifo len
	ser     serial.Serial
}

func (p *PL011) Init() error {
	if p.wr == nil {
		return errors.New("No Transmit method defined")
	}

	for i := 0; i < len(p.regs); i++ {
		switch i {
		case UARTDR:
			p.regs[i] = &regDR
			flen := 16
			if fifolen {
				flen = fifolen
			}
			p.regs[i].fifo = make([]u8, flen, flen<<1)

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

// This will be called from
func (p *PL011) SetInterruptLine(c chan bool) {

}

type regDR struct {
	fifo []u8
}

func (pl *PL011) write32(off uint32, val uint32) error {
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

func (pl *PL011) Write(off uint32, val uint32) error {
	return write32(off, val)
}

func (pl *PL011) read32(off uint32) (uint32, error) {

}

func (pl *PL011) read8(off uint32) (uint8, error) {
	val, err := read32(off)

	return uint8(val), err
}

func (pl *PL011) read16(off uint32) (uint16, error) {
	val, err := read32(off)

	return uint16(val), err
}

func (pl *PL011) Read(off uint32) (uint32 error) {
	return read32(off)
}
