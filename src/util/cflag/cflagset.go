// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
	Package cflag implements command-line flag parsing. but built on top of existing Flags
	cflag is a comma separated flag, but belongs to one flagSet.
	The notion of Flag, FlagSet is little different from "flag".

	Usage:
		import "cflag"
		p := cflag.NewFlagSet("perf", flag.ExitOnError)
		p.Add(cflag.Int("flagname", 1234, "help message for flagname"))
		p.Add(cflag.Int("anothername", 1, "Help message new"))
		flag.Var(p, "perf", "Usage definition, default value")

 			OR
		var flagvar int
		var ip = cflag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")

	at the command line,
	   <progname> -perf flagname=0
	   <progname> -perf flagname=0,anthername=123
*/

package cflag

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ErrHelp is the error returned if the flag -help is invoked but no such flag is defined.
var ErrHelp = errors.New("flag: help requested")

// -- bool Value
type boolValue bool

func newBoolValue(val bool, p *bool) *boolValue {
	*p = val
	return (*boolValue)(p)
}

func (b *boolValue) Set(s string) error {
	v, err := strconv.ParseBool(s)
	*b = boolValue(v)
	return err
}

func (b *boolValue) Get() interface{} { return bool(*b) }

func (b *boolValue) String() string { return fmt.Sprintf("%v", *b) }

func (b *boolValue) IsBoolFlag() bool { return true }

// optional interface to indicate boolean flags that can be
// supplied without "=value" text
type boolFlag interface {
	flag.Value
	IsBoolFlag() bool
}

// -- int Value
type intValue int

func newIntValue(val int, p *int) *intValue {
	*p = val
	return (*intValue)(p)
}

func (i *intValue) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = intValue(v)
	return err
}

func (i *intValue) Get() interface{} { return int(*i) }

func (i *intValue) String() string { return fmt.Sprintf("%v", *i) }

// -- int64 Value
type int64Value int64

func newInt64Value(val int64, p *int64) *int64Value {
	*p = val
	return (*int64Value)(p)
}

func (i *int64Value) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	*i = int64Value(v)
	return err
}

func (i *int64Value) Get() interface{} { return int64(*i) }

func (i *int64Value) String() string { return fmt.Sprintf("%v", *i) }

// -- uint Value
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

// -- uint64 Value
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

// -- string Value
type stringValue string

func newStringValue(val string, p *string) *stringValue {
	*p = val
	return (*stringValue)(p)
}

func (s *stringValue) Set(val string) error {
	*s = stringValue(val)
	return nil
}

func (s *stringValue) Get() interface{} { return string(*s) }

func (s *stringValue) String() string { return fmt.Sprintf("%s", *s) }

// -- float64 Value
type float64Value float64

func newFloat64Value(val float64, p *float64) *float64Value {
	*p = val
	return (*float64Value)(p)
}

func (f *float64Value) Set(s string) error {
	v, err := strconv.ParseFloat(s, 64)
	*f = float64Value(v)
	return err
}

func (f *float64Value) Get() interface{} { return float64(*f) }

func (f *float64Value) String() string { return fmt.Sprintf("%v", *f) }

// -- time.Duration Value
type durationValue time.Duration

func newDurationValue(val time.Duration, p *time.Duration) *durationValue {
	*p = val
	return (*durationValue)(p)
}

func (d *durationValue) Set(s string) error {
	v, err := time.ParseDuration(s)
	*d = durationValue(v)
	return err
}

func (d *durationValue) Get() interface{} { return time.Duration(*d) }

func (d *durationValue) String() string { return (*time.Duration)(d).String() }

// A FlagSet represents a set of defined flags.  The zero value of a FlagSet
// has no name and has ContinueOnError error handling.
type FlagSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler.
	Usage func()

	name          string
	parsed        bool
	actual        map[string]*flag.Flag
	formal        map[string]*flag.Flag
	args          []string // arguments after flags
	errorHandling flag.ErrorHandling
	output        io.Writer // nil means stderr; use out() accessor
}

// sortFlags returns the flags as a slice in lexicographical sorted order.
func sortFlags(flags map[string]*flag.Flag) []*flag.Flag {
	list := make(sort.StringSlice, len(flags))
	i := 0
	for _, f := range flags {
		list[i] = f.Name
		i++
	}
	list.Sort()
	result := make([]*flag.Flag, len(list))
	for i, name := range list {
		result[i] = flags[name]
	}
	return result
}

func (f *FlagSet) out() io.Writer {
	if f.output == nil {
		return os.Stderr
	}
	return f.output
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
func (f *FlagSet) SetOutput(output io.Writer) {
	f.output = output
}

// VisitAll visits the flags in lexicographical order, calling fn for each.
// It visits all flags, even those not set.
func (f *FlagSet) VisitAll(fn func(*flag.Flag)) {
	for _, flag := range sortFlags(f.formal) {
		fn(flag)
	}
}

// Visit visits the flags in lexicographical order, calling fn for each.
// It visits only those flags that have been set.
func (f *FlagSet) Visit(fn func(*flag.Flag)) {
	for _, flag := range sortFlags(f.actual) {
		fn(flag)
	}
}

// Lookup returns the Flag structure of the named flag, returning nil if none exists.
func (f *FlagSet) Lookup(name string) *flag.Flag {
	return f.formal[name]
}

// Set sets the value of the named flag.
func (f *FlagSet) SetName(name, value string) error {
	fl, ok := f.formal[name]
	if !ok {
		return fmt.Errorf("no such flag -%v", name)
	}
	err := fl.Value.Set(value)
	if err != nil {
		return err
	}
	if f.actual == nil {
		f.actual = make(map[string]*flag.Flag)
	}
	f.actual[name] = fl
	return nil
}

// PrintDefaults prints, to standard error unless configured
// otherwise, the default values of all defined flags in the set.
func (f *FlagSet) PrintDefaults() {
	f.VisitAll(func(flag *flag.Flag) {
		format := "  -%s=%s: %s\n"
		if _, ok := flag.Value.(*stringValue); ok {
			// put quotes on the value
			format = "  -%s=%q: %s\n"
		}
		fmt.Fprintf(f.out(), format, flag.Name, flag.DefValue, flag.Usage)
	})
}

// defaultUsage is the default function to print a usage message.
func defaultUsage(f *FlagSet) {
	if f.name == "" {
		fmt.Fprintf(f.out(), "Usage:\n")
	} else {
		fmt.Fprintf(f.out(), "Usage of %s:\n", f.name)
	}
	f.PrintDefaults()
}

// NOTE: Usage is not just defaultUsage(CommandLine)
// because it serves (via godoc flag Usage) as the example
// for how to write your own usage function.

// Usage prints to standard error a usage message documenting all defined command-line flags.
// The function is a variable that may be changed to point to a custom function.
var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	//PrintDefaults()
}

// NFlag returns the number of flags that have been set.
func (f *FlagSet) NFlag() int { return len(f.actual) }

// Arg returns the i'th argument.  Arg(0) is the first remaining argument
// after flags have been processed.
func (f *FlagSet) Arg(i int) string {
	if i < 0 || i >= len(f.args) {
		return ""
	}
	return f.args[i]
}

// NArg is the number of arguments remaining after flags have been processed.
func (f *FlagSet) NArg() int { return len(f.args) }

// Args returns the non-flag arguments.
func (f *FlagSet) Args() []string { return f.args }

// BoolVar defines a bool flag with specified name, default value, and usage string.
// The argument p points to a bool variable in which to store the value of the flag.
func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string) {
	f.AddFlag(Var(newBoolValue(value, p), name, usage))
}

func BoolVar(p *bool, name string, value bool, usage string) *flag.Flag {
	return Var(newBoolValue(value, p), name, usage)
}

// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func Bool(name string, value bool, usage string) *flag.Flag {
	p := new(bool)
	pv := newBoolValue(value, p)
	return Var(pv, name, usage)
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func (f *FlagSet) IntVar(p *int, name string, value int, usage string) {
	f.AddFlag(Var(newIntValue(value, p), name, usage))
}

// IntVar defines an int flag with specified name, default value, and usage string.
// The argument p points to an int variable in which to store the value of the flag.
func IntVar(p *int, name string, value int, usage string) *flag.Flag {
	return Var(newIntValue(value, p), name, usage)
}

// Int defines an int flag with specified name, default value, and usage string.
// The return value is the address of an int variable that stores the value of the flag.
func Int(name string, value int, usage string) *flag.Flag {
	p := new(int)
	pv := newIntValue(value, p)
	return Var(pv, name, usage)
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string) {
	f.AddFlag(Var(newInt64Value(value, p), name, usage))
}

// Int64Var defines an int64 flag with specified name, default value, and usage string.
// The argument p points to an int64 variable in which to store the value of the flag.
func Int64Var(p *int64, name string, value int64, usage string) *flag.Flag {
	return Var(newInt64Value(value, p), name, usage)
}

// Int64 defines an int64 flag with specified name, default value, and usage string.
// The return value is the address of an int64 variable that stores the value of the flag.
func Int64(name string, value int64, usage string) *flag.Flag {
	p := new(int64)
	pv := newInt64Value(value, p)
	return Var(pv, name, usage)
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint variable in which to store the value of the flag.
func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string) {
	f.AddFlag(Var(newUintValue(value, p), name, usage))
}

// UintVar defines a uint flag with specified name, default value, and usage string.
// The argument p points to a uint  variable in which to store the value of the flag.
func UintVar(p *uint, name string, value uint, usage string) *flag.Flag {
	return Var(newUintValue(value, p), name, usage)
}

// Uint defines a uint flag with specified name, default value, and usage string.
// The return value is the address of a uint  variable that stores the value of the flag.
func Uint(name string, value uint, usage string) *flag.Flag {
	p := new(uint)
	pv := newUintValue(value, p)
	return Var(pv, name, usage)
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string) {
	f.AddFlag(Var(newUint64Value(value, p), name, usage))
}

// Uint64Var defines a uint64 flag with specified name, default value, and usage string.
// The argument p points to a uint64 variable in which to store the value of the flag.
func Uint64Var(p *uint64, name string, value uint64, usage string) *flag.Flag {
	return Var(newUint64Value(value, p), name, usage)
}

// Uint64 defines a uint64 flag with specified name, default value, and usage string.
// The return value is the address of a uint64 variable that stores the value of the flag.
func Uint64(name string, value uint64, usage string) *flag.Flag {
	p := new(uint64)
	pv := newUint64Value(value, p)
	return Var(pv, name, usage)
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
	f.AddFlag(Var(newStringValue(value, p), name, usage))
}

// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
func StringVar(p *string, name string, value string, usage string) *flag.Flag {
	return Var(newStringValue(value, p), name, usage)
}

// String defines a string flag with specified name, default value, and usage string.
// The return value is the address of a string variable that stores the value of the flag.
func String(name string, value string, usage string) *flag.Flag {
	p := new(string)
	pv := newStringValue(value, p)
	return Var(pv, name, usage)
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string) {
	f.AddFlag(Var(newFloat64Value(value, p), name, usage))
}

// Float64Var defines a float64 flag with specified name, default value, and usage string.
// The argument p points to a float64 variable in which to store the value of the flag.
func Float64Var(p *float64, name string, value float64, usage string) *flag.Flag {
	return Var(newFloat64Value(value, p), name, usage)
}

// Float64 defines a float64 flag with specified name, default value, and usage string.
// The return value is the address of a float64 variable that stores the value of the flag.
func Float64(name string, value float64, usage string) *flag.Flag {
	p := new(float64)
	pv := newFloat64Value(value, p)
	return Var(pv, name, usage)
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	f.AddFlag(Var(newDurationValue(value, p), name, usage))
}

// DurationVar defines a time.Duration flag with specified name, default value, and usage string.
// The argument p points to a time.Duration variable in which to store the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func DurationVar(p *time.Duration, name string, value time.Duration,
	usage string) *flag.Flag {
	return Var(newDurationValue(value, p), name, usage)
}

// Duration defines a time.Duration flag with specified name, default value, and usage string.
// The return value is the address of a time.Duration variable that stores the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func Duration(name string, value time.Duration, usage string) *flag.Flag {
	p := new(time.Duration)
	pv := newDurationValue(value, p)
	return Var(pv, name, usage)
}

// AddFlag defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func (f *FlagSet) AddFlag(fl *flag.Flag) {
	// Remember the default value as a string; it won't change.
	_, alreadythere := f.formal[fl.Name]
	if alreadythere {
		var msg string
		if f.name == "" {
			msg = fmt.Sprintf("flag redefined: %s", fl.Name)
		} else {
			msg = fmt.Sprintf("%s flag redefined: %s", f.name, fl.Name)
		}
		fmt.Fprintln(f.out(), msg)
		panic(msg) // Happens only if flags are declared with identical names
	}
	if f.formal == nil {
		f.formal = make(map[string]*flag.Flag)
	}
	f.formal[fl.Name] = fl
}

// Var defines a flag with the specified name and usage string. The type and
// value of the flag are represented by the first argument, of type Value, which
// typically holds a user-defined implementation of Value. For instance, the
// caller could create a flag that turns a comma-separated string into a slice
// of strings by giving the slice the methods of Value; in particular, Set would
// decompose the comma-separated string into the slice.
func Var(value flag.Value, name string, usage string) *flag.Flag {
	return &flag.Flag{name, usage, value, value.String()}
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (f *FlagSet) failf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	fmt.Fprintln(f.out(), err)
	f.usage()
	return err
}

// usage calls the Usage method for the flag set, or the usage function if
// the flag set is CommandLine.
func (f *FlagSet) usage() {
	if f.Usage == nil {
		defaultUsage(f)
	} else {
		f.Usage()
	}
}

// parseOne parses one flag. It reports whether a flag was seen.
func (f *FlagSet) parseOne() (bool, error) {
	if len(f.args) == 0 {
		return false, nil
	}
	name := f.args[0]
	if len(name) == 0 || len(name) == 1 {
		return false, nil
	}

	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		return false, f.failf("bad flag syntax: %s", name)
	}

	// it's a flag. does it have an argument?
	f.args = f.args[1:]
	has_value := false
	value := ""
	for i := 1; i < len(name); i++ { // equals cannot be first
		if name[i] == '=' {
			value = name[i+1:]
			has_value = true
			name = name[0:i]
			break
		}
	}
	m := f.formal
	fl, alreadythere := m[name] // BUG
	if !alreadythere {
		if name == "help" || name == "h" { // special case for nice help message.
			f.usage()
			return false, ErrHelp
		}
		return false, f.failf("flag provided but not defined: -%s", name)
	}
	if fv, ok := fl.Value.(boolFlag); ok && fv.IsBoolFlag() { // special case: doesn't need an arg
		if has_value {
			if err := fv.Set(value); err != nil {
				return false, f.failf("invalid boolean value %q for -%s: %v", value, name, err)
			}
		} else {
			fv.Set("true")
		}
	} else {
		// It must have a value, which might be the next argument.
		if !has_value && len(f.args) > 0 {
			// value is the next arg
			has_value = true
			value, f.args = f.args[0], f.args[1:]
		}
		if !has_value {
			return false, f.failf("flag needs an argument: -%s", name)
		}
		if err := fl.Value.Set(value); err != nil {
			return false, f.failf("invalid value %q for flag -%s: %v", value, name, err)
		}
	}
	if f.actual == nil {
		f.actual = make(map[string]*flag.Flag)
	}
	f.actual[name] = fl
	return true, nil
}

// Parse parses flag definitions from the argument list, which should not
// include the command name.  Must be called after all flags in the FlagSet
// are defined and before flags are accessed by the program.
// The return value will be ErrHelp if -help was set but not defined.
func (f *FlagSet) Set(arguments string) error {
	f.parsed = true
	f.args = strings.Split(arguments, ",")
	for {
		seen, err := f.parseOne()
		if seen {
			continue
		}
		if err == nil {
			break
		}
		switch f.errorHandling {
		case flag.ContinueOnError:
			return err
		case flag.ExitOnError:
			os.Exit(2)
		case flag.PanicOnError:
			panic(err)
		}
	}
	return nil
}

// Parsed reports whether f.Parse has been called.
func (f *FlagSet) Parsed() bool {
	return f.parsed
}

func (f *FlagSet) Get() interface{} {
	return f.formal
}

func (f *FlagSet) String() string {
	return "Hello World"
}

// NewFlagSet returns a new, empty flag set with the specified name and
// error handling property.
func NewFlagSet(name string, errorHandling flag.ErrorHandling) *FlagSet {
	f := &FlagSet{
		name:          name,
		errorHandling: errorHandling,
	}
	return f
}

// Init sets the name and error handling property for a flag set.
// By default, the zero FlagSet uses an empty name and the
// ContinueOnError error handling policy.
func (f *FlagSet) Init(name string, errorHandling flag.ErrorHandling) {
	f.name = name
	f.errorHandling = errorHandling
}

// A MultiFlagSet is an option where the suboptions depend on current
// backend
type MultiFlagSet struct {
	name     string
	backends map[string]*FlagSet
	sep      string // Separator
	debug    bool
}

func NewMultiFlagSet(name string) *MultiFlagSet {
	return &MultiFlagSet{
		name:     name,
		backends: make(map[string]*FlagSet),
		sep:      ",",
	}
}

func (m *MultiFlagSet) AddOpt(b string, f *flag.Flag) {
	be, ok := m.backends[b]
	if !ok {
		be = NewFlagSet(b, flag.ExitOnError)
	}
	be.AddFlag(f)
	m.backends[b] = be
}

func (m *MultiFlagSet) Set(s string) error {
	i := strings.Index(s, ",")
	if i == -1 {
		return errors.New("Backend not specified")
	}
	o, ok := m.backends[s[:i]]
	if !ok {
		return fmt.Errorf("No such Backend registerd: %s %v", s[:i], o)
	}

	//o.Set(s[i+1:])
	return nil
}
