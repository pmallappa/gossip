package primecell

import (
	"dev"
)

const (
	PeriphID0 = 0xFE0 // PeriphID0 Register
	PeriphID1 = 0xFE4 // PeriphID1 Register
	PeriphID2 = 0xFE8 // PeriphID2 Register
	PeriphID3 = 0xFEC // PeriphID3 Register
	PCellID0  = 0xFF0 // PCellID0 Register
	PCellID1  = 0xFF4 // PCellID1 Register
	PCellID2  = 0xFF8 // PCellID2 Register
	PCellID3  = 0xFFC // PCellID3 Register
)

type PrimeCellID struct {
	idregs [8]dev.Register
}
