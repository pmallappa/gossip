package plat

import (
	"errors"
	"flag"
	"strconv"
)

import (
	"util"
)

var (
	platflags string
	smpflag   string
)


/* Try to process as much as possible, rest send to specific platform
for interpretation */
func parsePlatFlags() (map[string]string, error) {
	m, e := util.ParseFlagsSubst(platflags, "plat")
	if e != nil {
		goto out
	}
	for k, v := range m {
		switch k {
		case "mem":
			//p.SetMem(memParse(v))
			memSize, _ = util.ParseMem(v)
		case "?":
			var s string
			for i := range availplats {
				s += " vendor: " + availplats[i].GetInfo()["vendor"] +
					" model: " + availplats[i].GetInfo()["model"] + "\n"
			}
			e = errors.New(s)
		case "vendor":
			vendor = v
		case "model":
			model = v
		case "plat":
			platname = v
		default:
			continue
		}
		// if any cases returns non-nil
		if e != nil {
			goto out
		}
		// Delete the consumed options
		delete(m, k)
	}

	// Even at the end of this, we dont have a platform
out:
	return m, e
}

func parseSMPFlags() (int, error) {
	var smp, maxcpus, cores, threads, sockets uint64 = 1, 1, 1, 1, 1

	m, e := util.ParseFlagsSubst(smpflag, "smp")
	if e != nil {
		return 0, e
	}

	for k, v := range m {
		switch k {
		case "maxcpus":
			maxcpus, e = strconv.ParseUint(v, 0, 0)
		case "cores":
			cores, e = strconv.ParseUint(v, 0, 0)
		case "threads":
			threads, e = strconv.ParseUint(v, 0, 0)
		case "sockets":
			sockets, e = strconv.ParseUint(v, 0, 0)
		case "smp":
			smp, e = strconv.ParseUint(v, 0, 0)
		default:
		}

		delete(m, k)

		if e != nil {
			goto out
		}

	}

	// Suppress error untill we figure out the meaning of 'maxcpus'
	_ = maxcpus

	// smp is number of cores pers socket, number threads per core
	// and number of such sockets
	// smp = cores * threads * sockets

	// Need to compute the SMP options form what ever is given
	// If alone smp is given, using above equation calculate other values
	// sockets = 1
	// cores = smp / (sockets * threads)
	// threads = smp / (sockets * cores)

	// Recalculate
	if smp != sockets*cores*threads {
		if sockets > 1 {
			smp = sockets * cores * threads
		} else {
			cores = smp / (sockets * threads)
			threads = smp / (sockets * cores)
		}
	}

out:
	return int(smp), e
}

func ParseFlags() (map[string]string, error) {
	var m map[string]string
	var err error
	if m, err = parsePlatFlags(); err != nil {
		return nil, err
	}
	if nSMP, err = parseSMPFlags(); err != nil {
		return nil, err
	}

	return m, nil
}



func init() {
	util.PrintMe()
	availplats = make([]Platform, 0, 128)

	flag.StringVar(&platflags, "plat", "", "Platforms, type ? to list")
	flag.StringVar(&smpflag, "smp", "",
		"-smp n[,maxcpus=cpus][,cores=cores][,threads=threads][,sockets=sockets]")
}
