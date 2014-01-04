package main

import (
	"flag"
	"fmt"
	"strings"
	"util/cflag"
)

type Freq struct {
	c cflag.UnitsDec
}

func (f *Freq) Parse(s string) error {
	return nil
}
func (f *Freq) String() string {
	return fmt.Sprintf("%d\n", uint64(f.c))
}

func (f *Freq) Set(s string) error {
	slen := len(s)
	// Adjust to loose hz, if specified like 800Mhz
	if strings.HasSuffix(strings.ToLower(s), "hz") {
		s = s[:slen-2]
	}
	fmt.Printf("=====================================\n")
	return f.c.Parse(s)
}

// For logger we support option like
// If options supports multiple key=value, different separator can be used at each level
// first level is ',', second is ';', third is '!'
// -cpu freq=100MHz,log='level=WARNING;out=tcp:localhost:2000'

func main() {

	var f Freq
	var c cflag.CFlagSet

	flag.Var(&f, "freq", "CPU Freqency, accepts {K,M,G,k,m,g}Hz")
	flag.Var(&c, "cpu", "CPU Args")

	c.Add

	flag.Parse()
	fmt.Printf("%q\n", f)
}
