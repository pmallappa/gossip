package main

import (
	//"errors"
	"flag"
	"fmt"
	"os"
	//"strings"
)

import (
	"plat"
)

var (
	perf, debug   bool
	pstart, pstop int
	dstart, dstop int
	quiet, help   bool
)

func init() {
	flag.BoolVar(&quiet, "quiet", true, "As the name indicates, default true")

	// Performance
	flag.BoolVar(&perf, "perf", false, "enable perfomance measurement, summary printed at exit")
	flag.IntVar(&pstart, "pstart", -1, "[cycle] start performance measurement at cycle, summary printed at exit")
	flag.IntVar(&pstop, "pstop", -1, "[cycle] stop performance measurement at cycle, summary printed at exit")

	// Debug
	flag.BoolVar(&debug, "debug", false, "enables instruction tracing")
	flag.IntVar(&dstart, "dstart", -1, "[cycle] starts instruction tracing at [cycle]")
	flag.IntVar(&dstop, "dstop", -1, "[cycle] stops instruction tracing at [cycle]")

	// Help
	flag.BoolVar(&help, "help", false, "Display help")
	flag.BoolVar(&help, "h", false, "Display help")

}

func errExit1(errno error) {
	flag.Usage()
	os.Exit(1)
}

//
// Perfomance start and stop
//
func parsePerf() error {
	var start, stop int
	if perf {
		start = 0
		stop = 0
	}
	if pstart != -1 {
		start = pstart
	}
	if pstop != -1 {
		stop = pstop
	}
	pstop = stop
	pstart = start

	return nil
}

func parseLocal() error {
	var err error
	if err = parsePerf(); err != nil {
	}

	if err != nil {

	}

	return err
}

func parseFlags() {
	var err error = nil
	flag.Parse()

	if help {
		errExit1(nil)
	}

	if _, err = plat.ParseFlags(); err != nil {
		fmt.Printf("%v\n", err)
		errExit1(err)
	}
	if err = parseLocal(); err != nil {
		errExit1(err)
	}
}
