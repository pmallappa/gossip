// flag.go: CFlag is a comma separated list of flags
//          Supports extending via Parse() interface
package cflag

import (
	"flag"
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

type optType uint32

const (
	INT optType = iota
	UINT
	FLOAT
	STRING
	BOOL
	COMPLEX
	RUNE
	OTHER
)

type CFlag struct {
	name      string
	t         optType
	val       flag.Value
	dval      string // Default value
	shortname string // short form of operation, not needed
	desc      string // Description

	//TODO: Obsolete things, should be removed for good
	defval interface{}
	value  interface{}

	debug bool // For Testing only
}

func NewCFlag1(v flag.Value, name, desc, defval string, ot optType) *CFlag {
	c := &CFlag{
		val:  v,
		name: name,
		desc: desc,
		dval: defval,
		t:    ot,
	}
	c.val.Set(defval)
	return c
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
	return fmt.Sprintf("Name %s, Default: %s, Description:%s, Value:%v",
		cf.name, cf.dval, cf.desc, cf.val)
}

func (cf *CFlag) PrintDefaults() {
	if cf != nil && cf.defval != nil {
		fmt.Printf("Default:%q\n", cf.defval)
	}
}

func (cf *CFlag) Help() string {
	return fmt.Sprintf("\t%s\t %s (Default: %s)\n", cf.name, cf.desc, cf.val)
}

// cflag is comma separated key=value pairs
type cflagsetT struct {
	str   string
	cflag map[string]*CFlag
	sep   string // Separator used to delimit flags

	debug bool // For testing only
}

func New() *cflagsetT {
	return NewCFlagSet("")
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
func (cfs *cflagsetT) Set(str string) (e error) {
	cfs.str = str
	if cf, e := cfs.parse(); e != nil {
		cf.PrintDefaults()
	}
	return
}

func (cfs *cflagsetT) String() string {
	str := fmt.Sprintf("%s\n", cfs.str)
	for _, cf := range cfs.cflag {
		str += fmt.Sprintf("%q\n", cf)
	}
	return str
}

// -- END Value interface End

// -- BEGIN Getter interface
func (cfs *cflagsetT) Get() interface{} {
	return cfs
}

// -- END Getter

func (cfs *cflagsetT) GetOpt(str string) interface{} {
	return cfs.cflag[str].value
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
	if k.name != "" && k.name != split[0] {
		return fmt.Errorf("Unbelievable")
	}

	switch k.t {
	case INT:
		k.value, err = strconv.ParseInt(split[1], 0, 64)

	case UINT:
		k.value, err = strconv.ParseUint(split[1], 0, 64)

	case FLOAT:
		k.value, err = strconv.ParseFloat(split[1], 64)

	case BOOL:
		switch val := strings.ToLower(split[1]); val {
		case "yes", "true", "on", "1", "t":
			k.value = true
		case "no", "false", "off", "0", "f":
			k.value = false
		default:
			err = fmt.Errorf("Un-known value %s\n", val)
		}

	case STRING:
		k.value = split[1]

	default:
		err = k.val.Set(split[1])
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
		if cfs.debug {
			fmt.Printf("Parsing...%q\n", split[i])
		}
		if e = cfs.parseOne(split[i]); e != nil {
			return cfs.cflag[cf], e
		}
	}
	return
}

/*
* Add() function allows platform/devices to install default
* values, just in-case
 */
func (cfs *cflagsetT) Add(cf *CFlag) {
	cfs.cflag[cf.name] = cf
	if cfs.debug {
		fmt.Printf("Added %q\n", cf)
	}
}
