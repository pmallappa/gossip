package plat

import (
	"fmt"
	//"strings"
)

import (
	"bus"
	"cpu"
	//"dev"
	//"dev/mem"
	//"dev/serial"
	//"dev/net"
	"util/logng"
)

type PlatInfo struct {
	model   string
	vendor  string
	version string
}

func NewPlatInfo(m string, v string, ver string) *PlatInfo {
	return &PlatInfo{
		model:   m,
		vendor:  v,
		version: ver,
	}
}

type PlatRegister interface {
	Setup() Platform
}

type PlatInformer interface {
	GetInfo() map[string]string
	SetInfo(string, string, string)
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
	def      bool
	NumCores int // For Easy Access, its actully len(Cores)
	MemSize  uint64

	Cores []cpu.Cores

	//devices []dev.Devicer // An array of all devices on platform

	busMain bus.Bus
	// uart  []*serial.Serial

	// netdev []*net.Netdev
	// VGA: Some platforms like PC

	logger *logng.LoggerNG
}

type Platform interface {
	PlatInformer
	PlatController
	PlatDebugger
	PlatCustomizer
	PlatELFLoader
}

var (
	availplats  []Platform
	curPlatform Platform
	curPlat     Plat
	nSMP        int
)

var (
	memSize  uint64
	platname string
	vendor   string
	model    string
)

type PlatController interface {
	Init() error
	Start() error
	Stop() error
	ParseFlags() error
}

type PlatDebugger interface {
	Pause() error
	Resume() error
}

// Every Platform should implement this. for setting up
// its own things
type PlatCustomizer interface {
	PreSetup() error
	CustomSetup() error
	PostSetup() error
}

func (p *Plat) Finalize() error {
	// It is expected that All devices are added by actual platform
	// for _, d := range p.devices {
	// 	if e := d.Initialize(); e != nil {
	// 		fmt.Printf("Device %v did not initialize %v\n", d)
	// 	}
	//}
	return nil
}

func Register(p Platform) {
	availplats = append(availplats, p)
}

func NewPlat() *Plat {
	return &Plat{
		//PlatInfo: PlatInfo{model: model, vendor: vendor, version: "0.0"},
		logger: logng.NewLoggerNG(),
	}
}

func (p *Plat) Setup() error {
	if model == "" {
		return fmt.Errorf("Vendor not found, select a platform")
	}

	// search for right model
	for i := range availplats {
		if vendor == availplats[i].GetInfo()["vendor"] {

		}
	}

	return nil
}
