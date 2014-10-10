package main

import (
	"flag"
	"fmt"
	//"strings"
)
import (
	"../../cflag"
)

// For logger we support option like
// If options supports multiple key=value, different separator can be used at each level
// first level is ',', second is ';', third is '!'
// -cpu freq=100MHz,log='level=WARNING;out=tcp:localhost:2000'

var (
	quiet, help bool
)

func main() {
	perf := cflag.NewFlagSet("perf", flag.ExitOnError)
	for _, v := range []*flag.Flag{
		cflag.Int("start", -1, "[cycle] start performance measurement at cycle"),
		cflag.Int("stop", -1, "[cycle] stop performance measurement at cycle"),
	} {
		perf.AddFlag(v)
	}
	flag.Var(perf, "perf", "Default perf")

	// Debug
	debug := cflag.NewFlagSet("debug", flag.ExitOnError)
	for _, v := range []*flag.Flag{
		cflag.Int("start", -1, "[cycle] starts instruction tracing at [cycle]"),
		cflag.Int("stop", -1, "[cycle] stops instruction tracing at [cycle]"),
	} {
		debug.AddFlag(v)
	}

	flag.Var(debug, "debug", "Default debugs")

	// Help
	flag.BoolVar(&help, "help", false, "Display help")
	flag.BoolVar(&help, "h", false, "Display help")

	flag.BoolVar(&quiet, "quite", true, "As the Name says, be quite	")

	flag.Parse()

	var m = make(map[string]*flag.Flag)
	m = debug.Get().(map[string]*flag.Flag)
	for i, v := range m {
		fmt.Printf("i:%s v.Name:%s v:%q\n", i, v.Name, v.Value.String())
	}

	fmt.Printf("")
}
