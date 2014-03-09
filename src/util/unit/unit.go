package unit

import (
	"fmt"
	"strconv"
	"strings"
)

type Freq struct {
	b Binary
}

func (f *Freq) String() string {
	return fmt.Sprintf("%s", f.b.String())
}

func (f *Freq) Set(s string) error {
	slen := len(s)
	// Adjust to loose hz, if specified like 800Mhz
	if strings.HasSuffix(strings.ToLower(s), "hz") {
		s = s[:slen-2]
	}
	f.b.Set(s)

	return nil //f.c.Parse(s)
}

type Size struct {
	b Binary
}

func (f *Size) String() string {
	return fmt.Sprintf("%s", f.b.String())
}

func (f *Size) Set(s string) error {
	slen := len(s)
	// Adjust to loose hz, if specified like 800Mhz
	switch {
	case strings.HasSuffix(strings.ToLower(s), "i"):
		s = s[:slen-2]
	case strings.HasSuffix(strings.ToLower(s), "b"):
		s = s[:slen-1]
	}
	f.b.Set(s)

	return nil //f.c.Parse(s)
}

///
/// Special types to be parsed, eg memory size and freq
///
type (
	Binary  uint64
	Decimal uint64
)

// -- BEGIN Value interface
// MemSize will support MiB
func (b *Binary) Set(str string) (e error) {
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

	*b = Binary(mem * mult)

	return
}

func (b *Binary) String() string {
	str := ""
	val := uint64(*b)
	switch {
	case val > 1<<30:
		str = "K"
		val >>= 30
	case val > 1<<20:
		str = "M"
		val >>= 20
	case val > 1<<10:
		str = "G"
		val >>= 10
	}
	return fmt.Sprintf("%d%s", val, str)
}

// -- End Value interface

// Frequency will honour K/M/G or k/m/g
// to specify Kilo/Mega/Giga Hz.
// Eg: -cpu freq=800Mhz or freq=1000 or freq=200MHZ
func (d *Decimal) Set(s string) (e error) {
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

	*d = Decimal(dec * mult)
	return
}

func (d *Decimal) String() string {
	return fmt.Sprintf("%d", uint64(*d))
}
