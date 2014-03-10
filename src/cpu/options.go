package cpu

import (
	"flag"
	//"strings"
)

import (
	"util/cflag"
	"util/logng"
	"util/unit"
)

var cpuOpts = cflag.New()

var logger logng.LogNG

// For logger we support option like
// -cpu freq=100MHz,log='level=WARNING,out=tcp:localhost:2000'

func initCpuOpts() error {
	for _, v := range []*cflag.SubOption{
		cflag.NewSubOptionOther(&logger, "log", "Specify logging use <logdesc>",
			""),
		cflag.NewSubOptionOther(&unit.Freq{}, "freq",
			"CPU Freqency, accepts {K,M,G,k,m,g}Hz", "100MHz"),
	} {
		cpuOpts.Add(v)
	}

	flag.Var(cpuOpts, "cpu", "Various options for CPU")

	return nil
}

func GetOpt(s string) interface{} {
	return cpuOpts.GetSubOpt(s)
}
