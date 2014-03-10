package plat

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

import (
	"util"
	"util/cflag"
	"util/unit"
)

type smpFlags string

var smp int

func (s *smpFlags) String() string {
	return ""
}
func (s *smpFlags) Set(str string) error {
	var e error
	var tmp int64

	i := strings.Index(str, ",")

	if i == -1 {
		tmp, e = strconv.ParseInt(str, 0, 64)
	} else {
		tmp, e = strconv.ParseInt(str[:i-1], 0, 64)
	}
	if e != nil {
		return e
	}
	smp = int(tmp)
	smpdetails.Set(str[i+1:]) // remove the instance and proceed

	return nil
}

func parseSMP() (int, error) {
	var e error

	maxcpus := smpdetails.GetSubOpt("maxcpus").(int)
	cores := smpdetails.GetSubOpt("cores").(int)
	threads := smpdetails.GetSubOpt("threads").(int)
	sockets := smpdetails.GetSubOpt("sockets").(int)

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
	fmt.Printf("Cores:%d, Threads:%d sockets:%d\n", cores, threads, sockets)
	return int(smp), e
}

var (
	smpflags   smpFlags
	platflag   = cflag.New()
	smpdetails = cflag.New()
)

func initPlatOpts() {
	for _, v := range []*cflag.SubOption{
		cflag.NewSubOptionOther(&unit.Size{}, "mem", "Platform Memory",
			"128MiB"),
		cflag.NewSubOption("model", "Platform Model", ""),
		cflag.NewSubOption("vendor", "Platform Vendor", ""),
	} {
		platflag.Add(v)
	}

	flag.Var(platflag, "plat", "Platforms, type ? to list")
}

func initSMPOpts() {
	for _, v := range []*cflag.SubOption{
		cflag.NewSubOption("maxcpus", "Platform Memory", 1),
		cflag.NewSubOption("cores", "Platform Memory", 1),
		cflag.NewSubOption("threads", "Platform Memory", 1),
		cflag.NewSubOption("sockets", "Platform Memory", 1),
	} {
		smpdetails.Add(v)
	}

	flag.Var(&smpflags, "smp", "n[,maxcpus=cpus][,cores=cores][,threads=threads][,sockets=sockets]")
}

func init() {
	util.PrintMe()
	availplats = make([]Platform, 0, 128)

	initPlatOpts()
	initSMPOpts()
	initLdOpts()
}
