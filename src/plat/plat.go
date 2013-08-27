package plat

import (
	"errors"
	"flag"
	"fmt"

//"strings"
//"strconv"
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

	devices []dev.Devicer // An array of all devices on platform

	// uart  []serial.Serial

	// netdev []net.Netdev
	// VGA: Some platforms like PC
}

var availPlats []PlatInfo

var (
	memSize uint64
	vendor  string
	model   string
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

	flag.StringVar(&plat_opts, "plat", "", "Platforms, type ? to list")
	flag.StringVar(&smp, "smp", "",
		"-smp n[,maxcpus=cpus][,cores=cores][,threads=threads][,sockets=sockets]")

}

/* Try to process as much as possible, rest send to specific platform
for interpretation */
func ParseFlags() (map[string]string, error) {
	m, e := util.ParseFlags(plat_opts)
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
