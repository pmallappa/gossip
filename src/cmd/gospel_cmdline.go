package main

import (
	//"errors"
	"flag"
	//"fmt"
	"os"
	//"strings"
)

import (
	"cflag"
)

var (
	quiet, help bool
)

func init() {
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
}

func parseFlags() {
	var err error = nil
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(1)
	}
}
