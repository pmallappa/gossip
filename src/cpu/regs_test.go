package cpu

import (
	"testing"
)

func TestGpr(t *testing.T) {
	var reg Gpr
	reg.SetName("Myreg")
	if reg.Name != "Myreg" {
		t.Error("Name not updated")
	}

	reg.Set(0x03)
	if reg.Val() != 0x03 {
		t.Error("Value not found")
	}

}
