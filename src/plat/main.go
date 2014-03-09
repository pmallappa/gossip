package plat

import (
	"flag"
	//	"fmt"
	//"strconv"
)

import (
	"util"
	"util/cflag"
	"util/unit"
)

var (
	platflags string
	smpflag   string
)

func parseSMPFlags() (int, error) {
	var e error
	var smp, maxcpus, cores, threads, sockets uint64 = 1, 1, 1, 1, 1

	// Suppress error untill we figure out the meaning of 'maxcpus'
	_ = maxcpus

	// smp is number of cores pers socket, number threads per core
	// and number of such sockets
	// smp = cores * threads * sockets

	// Need to compute the SMP options form what ever is given
	// If alone smp is given, using above equation calculate other values
	// sockets = 1
	// cores = smp / (sockets * threads)
	// threads = smp / (sockets * cores)

	// Recalculate
	if smp != sockets*cores*threads {
		if sockets > 1 {
			smp = sockets * cores * threads
		} else {
			cores = smp / (sockets * threads)
			threads = smp / (sockets * cores)
		}
	}

	return int(smp), e
}

var (
	platFlags = cflag.New()
	smpFlags  = cflag.New()
)

func initPlatOpts() {
	for _, v := range []*cflag.CFlag{
		cflag.NewCFlag1(&unit.Size{}, "mem", "Platform Memory",
			"128MiB", cflag.OTHER),
		cflag.NewCFlag("model", "Platform Model", ""),
		cflag.NewCFlag("vendor", "Platform Vendor", ""),
	} {
		platFlags.Add(v)
	}

	flag.Var(platFlags, "plat", "Platforms, type ? to list")
}

func initSMPOpts() {
	for _, v := range []*cflag.CFlag{
		cflag.NewCFlag("maxcpus", "Platform Memory", 1),
		cflag.NewCFlag("cores", "Platform Memory", 1),
		cflag.NewCFlag("threads", "Platform Memory", 1),
		cflag.NewCFlag("sockets", "Platform Memory", 1),
	} {
		smpFlags.Add(v)
	}
	flag.Var(smpFlags, "smp", "SMP")
}

func init() {
	util.PrintMe()
	availplats = make([]Platform, 0, 128)

	initPlatOpts()
	initSMPOpts()
	initLdOpts()
}
