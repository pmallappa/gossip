package arm

import (
	//"fmt"
)

import (
//	"cpu"
	"cpu/arm"
)

var supported_cpus = []string {
	"926EJS",
	"1176JZFS",
	"11MPCore",
	"CortexA9",
	"CortexA15",
	"CortexA53",
	"CortexA57",
}

func pbInit() {
	println("Init plat/arm/realviewPB")
	for c := range supported_cpus {
		arm.RegisterCpu(supported_cpus[c])
	}
}
