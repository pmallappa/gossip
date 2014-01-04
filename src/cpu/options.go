package cpu

import (
	"flag"
	"strings"
)

import (
	"util"
	"util/cflag"
	"util/logng"
)

var cpuOpts = cflag.NewFlagSet("cpu")

var logger = logng.LogNG

// For logger we support option like
// -cpu freq=100MHz,log='level=WARNING,out=tcp:localhost:2000'

func initOpts() error {
	for _, v := range []*cflag.CFlag{
		cflag.NewCFlag("log", "Specify logging use <logdesc>", &logger),
		cflag.NewCFlag("freq", "CPU Freqency, accepts {K,M,G,k,m,g}Hz",
			cflag.UnitsDec(100*1024*1024)),
	} {
		cpuOpts.Add(v)
	}

	flag.Var(&cpuOpts, "cpu", "Various options for CPU")
}

func GetOpt(s string) interface{} {
	return cpuOpts.GetOpt(s)
}
