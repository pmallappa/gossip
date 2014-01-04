// flag.go: CFlag is a comma separated list of flags
//          Supports extending via Parse() interface
package cflag

import (
	//"flag"
	"fmt"
	"strconv"
	"strings"
)

type CFlagParser interface {
	Parse(string) error
}
type OptGetter interface {
	GetOpt(string) interface{}
}
type CFlagHelper interface {
	Help() string
}

// cflag is comma separated key=value pairs
type cflagsetT struct {
	str   string
	cflag map[string]*CFlag
	sep   byte // Separator used to delimit flags

	debug bool // For testing only
}

func NewCFlagSet(s string) *cflagsetT {
	return &cflagsetT{
		str:   s,
		sep:   ",",
		cflag: make(map[string]*CFlag),
	}
}

func (cfs *cflagsetT) PrintDefaults() {
	for _, v := range cfs.cflag {
		v.PrintDefaults()
	}
}

// -- BEGIN Value interface
func (cfs *cflagsetT) Set(str string) error {
	cfs.str = str
	if e = cfs.parse(); e != nil {
		cfs.PrintDefaults()
	}

}

func (cfs *cflagsetT) String() string {
	return cfs.str
}

// -- END Value interface End

// -- BEGIN Getter interface
func (cfs *cflagSetT) Get() interface{} {
	return cfs
}

// -- END Getter

func (cfs *cflagsetT) GetOpt(str string) interface{} {
	return cfs.cflag[str].value
}

type CFlag struct {
	name      string
	defval    interface{}
	shortname string // short form of operation, not needed
	desc      string // Description
	value     interface{}

	debug bool // For Testing only
}

func NewCFlag(name, desc string, def interface{}) *CFlag {
	return &CFlag{
		name:   name,
		desc:   desc,
		defval: def,
		value:  def,
	}
}

func (cf *CFlag) String() string {
	return fmt.Sprintf("Name %s, Default: %v, Description:%s, Value:%v",
		cf.name, cf.defval, cf.desc, cf.value)
}

func (cfs *cflagsetT) parseOne(s string) (err error) {

	split := strings.SplitN(s, "=", 2)

	if len(split) != 2 {
		return
	}
	k := cfs.cflag[split[0]]

	if k == nil {
		return fmt.Errorf("Not Found")
	}

	// Confirmation
	if k.name != split[0] {
		return fmt.Errorf("Unbelievable")
	}

	switch k.defval.(type) {
	case int8, int16, int32, int64, int:
		k.value, err = strconv.ParseInt(split[1], 0, 64)

	case uint8, uint16, uint32, uint64, uint:
		k.value, err = strconv.ParseUint(split[1], 0, 64)

	case float32, float64:
		k.value, err = strconv.ParseFloat(split[1], 64)

	case bool:
		switch val := strings.ToLower(split[1]); val {
		case "yes", "true", "on", "1", "t":
			k.value = true
		case "no", "false", "off", "0", "f":
			k.value = false
		default:
			err = fmt.Errorf("Unfavourable value %s\n", split[1])
		}

	case string:
		k.value = split[1]

	case UnitsBin:
		b := k.value.(UnitsBin)
		if err = b.Parse(split[1]); err != nil {
			k.value = b
		}

	case UnitsDec:
		d := k.value.(UnitsDec)
		if err = d.Parse(split[1]); err != nil {
			k.value = d
		}

	default:
		err = k.value.Set(split[1])
	}

	if k.debug && err == nil {
		fmt.Printf("Parsed %q\n", k)
	}

	return
}

// Real parse function, for each of key=value pairs,
func (cfs *cflagsetT) parse() (c *CFlag, e error) {
	str := cfs.str
	split := strings.Split(str, cfs.sep)

	for i, cf := range split {
		if e = cfs.parseOne(split[i]); e != nil {
			return cfs.cflag[cf], e
		}
	}
	return
}

func (cfs *cflagsetT) Add(cf *CFlag) {
	cfs.cflag[cf.name] = cf
	if cfs.debug {
		fmt.Printf("Added %q\n", cf)
	}
}

///
/// Special types to be parsed, eg memory size and freq
///
type (
	UnitsBin uint64
	UnitsDec uint64
)

// -- BEGIN Value interface
// MemSize will support MiB
func (b *UnitsBin) Set(str string) error {
	var mult, mem uint64

	strlen := len(str)

	switch str[strlen-1] {
	case 'k', 'K':
		mult = 1 << 10 // 2^10
	case 'm', 'M':
		mult = 1 << 20 // 2^20
	case 'g', 'G':
		mult = 1 << 30 // 2^30
	default:
		mult = 1
	}

	if mult != 1 {
		str = str[:strlen-1]
	}

	// This cannot be a 'int' rather be a 'uint'
	if mem, e = strconv.ParseUint(str, 0, 64); e != nil {
		println(e)
		return e
	}

	*b = UnitsBin(mem * mult)

	return
}

func (b *UnitsBin) String() string {
	return fmt.Sprintf("%d", uint64(*b))
}

// -- End Value interface

// Frequency will honour K/M/G or k/m/g
// to specify Kilo/Mega/Giga Hz.
// Eg: -cpu freq=800Mhz or freq=1000 or freq=200MHZ
func (d *UnitsDec) Set() error {
	var mult, dec uint64

	slen := len(s)

	switch s[slen-1] {
	case 'k', 'K':
		mult = 1e3 // 10^3
	case 'm', 'M':
		mult = 1e6 // 10^6
	case 'g', 'G':
		mult = 1e9 // 10^9
	default:
		mult = 1
	}

	// Adjust the string to loose last 'K/k/M/m/G/g'
	if mult != 1 {
		s = s[:slen-1]
	}
	// Though Uint32 is sufficient as processor freq are not more than 4GHz
	if dec, e = strconv.ParseUint(s, 0, 64); e != nil {
		return
	}

	*d = UnitsDec(dec * mult)
	return
}

func (d *UnitsDec) String() string {
	return fmt.Sprintf("%d", uint64(*d))
}
