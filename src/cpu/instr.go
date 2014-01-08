package cpu

// Defines all instruction related interface and structures

type instrT struct {
	length uint8
}

type Instr interface {
	String() string
	GetLen() uint
}
