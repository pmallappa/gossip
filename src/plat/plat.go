package plat

import (
	"errors"
	"flag"
	"fmt"
	//"strings"
	"strconv"
)

import (
	//	"bus"
	"cpu"
	"dev"
	//	"dev/mem"
	//	"dev/serial"
	//	"dev/net"
	"util"
)

var (
	platflags string
	smpflag   string
)

type PlatInfo struct {
	model   string
	vendor  string
	version string
}

func (p *PlatInfo) GetInfo() map[string]string {
	return map[string]string{"model": p.model, "vendor": p.vendor, "version": p.version}
}

func (p *PlatInfo) SetInfo(model, vendor, version string) {
	p.model = model
	p.vendor = vendor
	p.version = version
}

type Plat struct {
	PlatInfo
	Cores    []cpu.Cores
	NumCores int // For Easy Access, its actully len(Cores)
	MemSize  uint64

	devices []dev.Devicer // An array of all devices on platform

	// uart  []serial.Serial

	// netdev []net.Netdev
	// VGA: Some platforms like PC
}

var availPlats []PlatInfo

var (
	memSize  uint64
	platname string
	vendor   string
	model    string
)

type Platformer interface {
	Init() error
	Start()
	//ParseFlags() error
}

type PlatDebugger interface {
	Pause()
	Resume()
}

func (p *Plat) Finalize() error {
	// It is expected that All devices are added by actual platform
	for _, d := range p.devices {
		if e := d.Initialize(); e != nil {
			fmt.Printf("Device %v did not initialize %v\n", d)
		}
	}
	return nil
}

func NewPlat() *Plat {
	return &Plat{
		PlatInfo: PlatInfo{model: model, vendor: vendor, version: "0.0"},
		MemSize:  memSize,
	}
}

func init() {
	util.PrintMe()
	availPlats = make([]PlatInfo, 16)

	flag.StringVar(&platflags, "plat", "", "Platforms, type ? to list")
	flag.StringVar(&smpflag, "smp", "",
		"-smp n[,maxcpus=cpus][,cores=cores][,threads=threads][,sockets=sockets]")
}

/* Try to process as much as possible, rest send to specific platform
for interpretation */
func ParsePlatFlags() (map[string]string, error) {
	m, e := util.ParseFlagsSubst(platflags, "plat")
	if e != nil {
		goto out
	}
	for k, v := range m {
		switch k {
		case "mem":
			//p.SetMem(memParse(v))
			memSize, _ = util.ParseMem(v)
		case "?":
			var s string
			for i := range availPlats {
				s += " vendor: " + availPlats[i].vendor + " model: " +
					availPlats[i].model + "\n"
			}
			e = errors.New(s)
		case "vendor":
			vendor = v
		case "model":
			model = v
		case "plat":
			platname = v
		default:
			continue
		}
		// if any cases returns non-nil
		if e != nil {
			goto out
		}
		// Delete the consumed options
		delete(m, k)
	}
out:
	return m, e
}

func ParseSMPFlags() (int, error) {
	var smp, maxcpus, cores, threads, sockets uint64 = 1, 1, 1, 1, 1

	m, e := util.ParseFlagsSubst(smpflag, "smp")
	for k, v := range m {
		switch k {
		case "maxcpus":
			maxcpus, e = strconv.ParseUint(v, 0, 0)
		case "cores":
			cores, e = strconv.ParseUint(v, 0, 0)
		case "threads":
			threads, e = strconv.ParseUint(v, 0, 0)
		case "sockets":
			sockets, e = strconv.ParseUint(v, 0, 0)
		case "smp":
			smp, e = strconv.ParseUint(v, 0, 0)
		default:
			fmt.Printf("Dont understand options", k, v)
			continue
		}

		if e != nil {
			goto out
		}

	}

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

out:
	return int(smp), e
}

func ParseFlags() (map[string]string, error) {
	// TODO
	return nil, nil
}
