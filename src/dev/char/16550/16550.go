package 16550

import()

type uart16550 struct {
	hr uint8 // Holding register (transmit/recieve)
	ier uint8 // Interrup Enable
	isr uint8 // Interrup Status
	fcr uint8 // Fifo Control
	lcr uint8 // Line control
	mcr uint8 // Modem Control 
	lsr uint8 // Line Control register
	msr uint8 // Modem status register

	fifo[] uint8
}


func (u *uart16550) Read() (uint64, error) {
		return 0, nil
}

func (u *uart16550) Read16() (uint16, error) {
		return 0, nil
}

func (u *uart16550) Read32() (uint32, error) {
		return 0, nil
}

func (u *uart16550) Read64() (uint64, error) {
	return 0, nil
}

func (u *uart16550) Write(uint64, uint64) error {
		return 0, nil
}

func (u *uart16550) Write16(uint64, uint16) error {
		return 0, nil
}

func (u *uart16550) Write32(uint64, uint32) error {
		return 0, nil
}

func (u *uart16550) Write64(uint64, uint64) error {
		return 0, nil
}

