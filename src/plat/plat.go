package plat

import (
	"errors"
	"flag"

//	"fmt"
//"strings"
//"strconv"
)

import (
	"cpu"
	"dev"
	"dev/mem"
	//	"dev/serial"
	//	"dev/net"
	"util"
)

var (
	plat_opts string
	smp       string
)

type PlatInfo struct {
	model   string
	vendor  string
	version string
}

func (p *PlatInfo) GetInfo() map[string]string {
	return map[string]string{"model": p.model, "vendor": p.vendor, "version": p.version}
}

func (p *PlatInfo) SetInfo(info *PlatInfo) {
	p.model = info.model
	p.vendor = info.vendor
	p.version = info.version
}

type Plat struct {
	PlatInfo
	Cores    []cpu.Cores
	NumCores int // For Easy Access, its actully len(Cores)
	MemSize  uint64

	devices []dev.Device // An array of all devices on platform

	// uart  []serial.Serial

	// netdev []net.Netdev
	// VGA: Some platforms like PC
}

var availablePlatforms []PlatInfo

var (
	memSize uint64
	vendor  string
	model   string
)

type Platform interface {
	Init() error
	Start()
	//ParseFlags() error
}

func (p *Plat) Finalize() error {
	// It is expected that All devices are added by actual platform
	memory = mem.Newmem(p.MemSize)
	for device := range p.devices {
		e := device.Initialize()
		if e != nil {
			fmt.Printf("Device %v did not initialize %v\n", device)
		}
	}
}

func NewPlat() *Plat {
	return Plat{
		PlatInfo{vendor: vendor, model: model},
		MemSize: memSize,
	}

	return p
}

func init() {
	util.PrintMe()
	availablePlatforms = make([]PlatInfo, 16)

	flag.StringVar(&plat_opts, "plat", "", "Platforms, type ? to list")
	flag.StringVar(&smp, "smp", "",
		"-smp n[,maxcpus=cpus][,cores=cores][,threads=threads][,sockets=sockets]")

}

/* Try to process as much as possible, rest send to specific platform
for interpretation */
func ParseFlags() (map[string]string, error) {
	m, e := util.ParseFlags(plat_opts)
	for k, v := range m {
		switch {
		case k == "mem":
			//p.SetMem(memParse(v))
			memSize, _ := util.ParseMem(v)
		case k == "?":
			var s string
			for i := range availablePlatforms {
				s += " vendor: " + availablePlatforms[i].vendor + " model: " +
					availablePlatforms[i].model + "\n"
			}
			e = errors.New(s)
		case "vendor":
			vendor = v
		case "model":
			model = v
		default:
			continue
		}
		// if any cases returns non-nil
		if e != nil {
			return nil, e
		}
		// Delete the consumed options
		delete(m, k)
	}

	return m, e
}
