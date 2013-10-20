package cpu

import(
"testing"
)

func (testing *T) TestGpr() {
	var reg Gpr
	reg.SetName("Myreg")
	if reg.Name != "Myreg" {
		T.Fail("Name not updated")
	}

}