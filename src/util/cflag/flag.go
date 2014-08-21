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
	BUILTIN
	OTHER
)

type SubOption struct {
	name      string
	t         optType
	val       flag.Value
	defval    string // Default value
	shortname string // short form of operation, not needed
	desc      string // Description

	value interface{}

	debug bool // For Testing only
}

func NewSubOptionOther(v flag.Value, name, desc, defval string) *SubOption {
	c := &SubOption{
		val:    v,
		name:   name,
		desc:   desc,
		defval: defval,
		t:      OTHER,
	}
	c.val.Set(defval)
	return c
}

func NewSubOption(name, desc string, def interface{}) *SubOption {
	return &SubOption{
		name: name,
		desc: desc,
		//defval: def,
		value: def,
		t:     BUILTIN,
	}
}

func (cf *SubOption) String() string {
	return fmt.Sprintf("Name %s, Default: %s, Description:%s, Value:%v",
		cf.name, cf.defval, cf.desc, cf.val)
}

func (cf *SubOption) PrintDefaults() {
	fmt.Printf("Default:%q\n", cf.defval)
}

func (cf *SubOption) Help() string {
	return fmt.Sprintf("\t%s\t %s (Default: %s)\n", cf.name, cf.desc, cf.val)
}

// cflag is comma separated key=value pairs
type OptionT struct {
	str     string
	subopts map[string]*SubOption
	sep     string // Separator used to delimit flags

	debug bool // For testing only
}

func New() *OptionT {
	return NewOption("")
}

func NewOption(s string) *OptionT {
	return &OptionT{
		str:     s,
		sep:     ",",
		subopts: make(map[string]*SubOption),
	}
}

func (cfs *OptionT) PrintDefaults() {
	for _, v := range cfs.subopts {
		v.PrintDefaults()
	}
}

// -- BEGIN Value interface
func (cfs *OptionT) Set(str string) (e error) {
	cfs.str = str
	if cf, e := cfs.parse(); e != nil {
		cf.PrintDefaults()
	}
	return
}

func (cfs *OptionT) String() string {
	str := fmt.Sprintf("%s\n", cfs.str)
	for _, cf := range cfs.subopts {
		str += fmt.Sprintf("%q\n", cf)
	}
	return str
}

// -- END Value interface End

// -- BEGIN Getter interface
func (cfs *OptionT) Get() interface{} {
	return cfs
}

// -- END Getter

func (cfs *OptionT) GetSubOpt(str string) interface{} {
	return cfs.subopts[str].value
}

func (cfs *OptionT) parseOne(s string) (err error) {

	split := strings.SplitN(s, "=", 2)

	if len(split) != 2 {
		return
	}
	k := cfs.subopts[split[0]]

	if k == nil {
		return fmt.Errorf("Not Found")
	}

	// Confirmation
	if k.name != "" && k.name != split[0] {
		return fmt.Errorf("Unbelievable")
	}

	if k.t == BUILTIN {
		switch k.value.(type) {
		case int, int8, int16, int32, int64:
			k.value, err = strconv.ParseInt(split[1], 0, 64)

		case uint, uint8, uint16, uint32, uint64:
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
				err = fmt.Errorf("Un-known value %s\n", val)
			}

		case string:
			k.value = split[1]
		}
	} else {
		err = k.val.Set(split[1])
	}

	if k.debug && err == nil {
		fmt.Printf("Parsed %q\n", k)
	}

	return
}

// Real parse function, for each of key=value pairs,
func (cfs *OptionT) parse() (c *SubOption, e error) {
	str := cfs.str
	split := strings.Split(str, cfs.sep)

	for i, cf := range split {
		if cfs.debug {
			fmt.Printf("Parsing...%q\n", split[i])
		}
		if e = cfs.parseOne(split[i]); e != nil {
			return cfs.subopts[cf], e
		}
	}
	return
}

/*
* Add() function allows platform/devices to install default
* values, just in-case
 */
func (cfs *OptionT) Add(cf *SubOption) {
	cfs.subopts[cf.name] = cf
	if cfs.debug {
		fmt.Printf("Added %q\n", cf)
	}
}
