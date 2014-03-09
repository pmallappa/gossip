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

func initOpts() error {
	for _, v := range []*cflag.CFlag{
		cflag.NewCFlag1(&logger, "log", "Specify logging use <logdesc>",
			"", cflag.OTHER),
		cflag.NewCFlag1(&unit.Freq{}, "freq",
			"CPU Freqency, accepts {K,M,G,k,m,g}Hz", "100MHz",
			cflag.OTHER),
	} {
		cpuOpts.Add(v)
	}

	flag.Var(cpuOpts, "cpu", "Various options for CPU")

	return nil
}

func GetOpt(s string) interface{} {
	return cpuOpts.GetOpt(s)
}
