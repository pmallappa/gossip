package 16550

import()

import(
	"bus"
)

type uart16550 struct {
	hr uint8 // Holding register (transmit/recieve)
	ier uint8 // Interrupt Enable
	isr uint8 // Interrupt Status
	fcr uint8 // Fifo Control
	lcr uint8 // Line control
	mcr uint8 // Modem Control 
	lsr uint8 // Line Control register
	msr uint8 // Modem status register

	rd bus.Reader
	rw bus.Writer

	rfifo[] uint8 // Read Fifo
	wfifo[] uint8 // Write Fifo
}


func (u *uart16550) Read(addr uint64) (uint64, error) {
		return 0, nil
}

func (u *uart16550) Read16(addr uint64) (uint16, error) {
		return 0, nil
}

func (u *uart16550) Read32(addr uint64) (uint32, error) {
		return 0, nil
}

func (u *uart16550) Read64(addr uint64) (uint64, error) {
	return 0, nil
}

func (u *uart16550) Write(addr, val uint64) error {
		return 0, nil
}

func (u *uart16550) Write16(addr, val uint16) error {
		return 0, nil
}

func (u *uart16550) Write32(addr, val uint32) error {
		return 0, nil
}

func (u *uart16550) Write64(addr, val, uint64) error {
		return 0, nil
}

