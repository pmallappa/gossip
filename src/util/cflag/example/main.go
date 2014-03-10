package main

import (
	"flag"
	"fmt"
	"strings"
)
import (
	"util"
	"util/cflag"
	"util/unit"
)

type Freq struct {
	c unit.Decimal
}

func (f *Freq) Parse(s string) error {
	return nil
}

func (f *Freq) String() string {
	return fmt.Sprintf("%s", f.c.String())
}

func (f *Freq) Set(s string) error {
	slen := len(s)
	// Adjust to loose hz, if specified like 800Mhz
	if strings.HasSuffix(strings.ToLower(s), "hz") {
		s = s[:slen-2]
	}
	f.c.Set(s)
	fmt.Printf("============== Setting to :%d\n", f.c)
	return nil //f.c.Parse(s)
}

// For logger we support option like
// If options supports multiple key=value, different separator can be used at each level
// first level is ',', second is ';', third is '!'
// -cpu freq=100MHz,log='level=WARNING;out=tcp:localhost:2000'

func main() {

	//var c cflag.CFlagSet
	c := cflag.New()
	//flag.Var(c, "cpu", "CPU Freqency, accepts {K,M,G,k,m,g}Hz")

	c.Add(cflag.NewSubOptionOther(&Freq{}, "freq",
		"CPU Freqency, accepts {K,M,G,k,m,g}Hz)",
		"100MHZ"))

	flag.Var(c, "cpu", "accepts things")

	flag.Parse()
	fmt.Printf("%s\n", c)
}
