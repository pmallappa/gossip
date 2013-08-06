package cpu

import (
	"errors"
	"log"
)
import (
	"util"
)

// Static Declaration to pass values to new CPU's
var cpu CpuInfo

func ParseFlags() (map[string]string, error) {
	var e error
	m, e := util.ParseFlags(cpu_opts)
	for k, v := range m {
		switch k {
		case "freq":
			cpu.freq, e = util.ParseFreq(v)
		case "log":
			var logger *log.Logger
			if logger, e = util.ParseLogger(v); e != nil {
				return nil, e
			}
			//cpu.SetLogger(logger)
			logger = logger
		case "?":
			var s string
			for i := range availableCpu {
				s += " vendor: " + availableCpu[i].vendor + " model: " +
					availableCpu[i].model + "\n"
			}
			e = errors.New(s)
		default:
			continue // Skip this option, chip specific; may be
		}
		// if any cases returns non-nil
		if e != nil {
			return nil, e
		}
		// Delete the consumed options
		delete(m, k)
	}
	return m, e
}
