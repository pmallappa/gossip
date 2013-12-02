package cpu

import (
	"testing"
)

func TestGpr(t *testing.T) {
	var reg Gpr
	reg.SetName("Myreg")
	if reg.Name() != "Myreg" {
		t.Error("Name not updated")
	}

	reg.Set(0x03)
	if reg.Val() != 0x03 {
		t.Error("Value not found")
	}
}

func TestSpecl(t *testing.T) {
	var sreg SpclReg

	sreg.SetName("Mysplreg")
	if sreg.Name() != "Mysplreg" {
		t.Error("Name not updated")
	}

	sreg.SetReserved(0x101010101010, true)
	sreg.Set(0x010101010101)
	if sreg.Val() != 0x111111111111 {
		t.Error("Value incorrect expected %x, got %x", 0x111111111111,
			sreg.Val())
	}

	// Reset reserved ones,
	sreg.SetReserved(0x0, true)
	sreg.SetReserved(0x000fff, false)
	sreg.Set(0xffffff)
	if sreg.Val() != 0xfff000 {
		t.Error("Value incorrect expected %x, got %x", 0xfff000,
			sreg.Val())
	}
}
