package cflag

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type formatType uint8

const (
	fmtBin formatType = iota
	fmtOct
	fmtDec
	fmtHex
)

type Opt struct {
	flag.Flag
	fmtT formatType
}

type Opts struct {
	opt   map[string]*Opt
	sep   string
	debug bool
}

func NewOpts() *Opts {
	c := &Opts{
		//backend: b,
		sep: ",",
		opt: make(map[string]*Opt),
	}
	return c
}

func (o *Opts) Add(f *Opt) {
	o.opt[f.Name] = f
}

func (o *Opts) setOne(s string) error {
	split := strings.Split(s, "=")
	key := split[0]
	val := ""
	if len(split) == 1 {
		// Probably a boolean variable
		val = "on"
	} else {
		val = split[1]
	}

	opt, ok := o.opt[key]

	if !ok || opt == nil || opt.Value == nil {
		return errors.New("No such option")
	}

	return opt.Value.Set(val)
}

func (o *Opts) Set(s string) error {
	split := strings.Split(s, o.sep)
	for _, ss := range split {
		if e := o.setOne(ss); e != nil {
			return e
		}
	}
	return nil
}

func (o *Opts) Usage() {

}

func (o *Opts) String() string {
	s := ""
	for _, opt := range o.opt {
		s += opt.Name + ": " + opt.Value.String()
		s += "\t"
	}
	return s
}

// A MultiOpts is an option where the suboptions depend on current
// backend
type MultiOpts struct {
	backends map[string]*Opts
	sep      string // Separator
	debug    bool
}

func NewMultiOpts() *MultiOpts {
	return &MultiOpts{
		backends: make(map[string]*Opts),
		sep:      ",",
	}
}

func (m *MultiOpts) AddOpt(b string, f *Opt) {
	be, ok := m.backends[b]
	if !ok {
		be = NewOpts()
	}
	be.Add(f)
	m.backends[b] = be
}

func (m *MultiOpts) Set(s string) error {
	i := strings.Index(s, ",")
	if i == -1 {
		return errors.New("Backend not specified")
	}
	o, ok := m.backends[s[:i]]
	if !ok {
		return fmt.Errorf("No such Backend registerd: %s", s[:i])
	}

	o.Set(s[i+1:])
	return nil
}

func (m *MultiOpts) String() string {
	return ""
}

func (m *MultiOpts) Stringize() string {
	s := ""
	for k, _ := range m.backends {
		s += k
		//s += "\t" + v.Stringize()
		s += "\n"
	}
	return s
}

func Var(value flag.Value, name, usage string) *Opt {
	return &Opt{
		flag.Flag{name, usage, value, value.String()},
		fmtBin,
	}
}

// -- Bool

// -- String

// -- Uint
type uintValue uint

func newUintValue(val uint, p *uint) *uintValue {
	*p = val
	return (*uintValue)(p)
}

func (i *uintValue) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uintValue(v)
	return err
}

func (i *uintValue) Get() interface{} { return uint(*i) }

func (i *uintValue) String() string { return fmt.Sprintf("%v", *i) }

func Uint(name string, value uint, usage string) *Opt {
	p := new(uint)
	pv := newUintValue(value, p)
	return Var(pv, name, usage)
}

// -- Uint64
type uint64Value uint64

func newUint64Value(val uint64, p *uint64) *uint64Value {
	*p = val
	return (*uint64Value)(p)
}

func (i *uint64Value) Set(s string) error {
	v, err := strconv.ParseUint(s, 0, 64)
	*i = uint64Value(v)
	return err
}

func (i *uint64Value) Get() interface{} { return uint64(*i) }

func (i *uint64Value) String() string { return fmt.Sprintf("%v", *i) }

func Uint64(name string, value uint64, usage string) *Opt {
	p := new(uint64)
	pv := newUint64Value(value, p)
	return Var(pv, name, usage)
}

// -- Int

// -- Int64
