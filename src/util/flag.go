package util

import (
	//"fmt"
	//"errors"
	//"io"
	//"log"
	//"os"
	"strconv"
	"strings"
)

import (
	//"util/logger"
)



func _parseflags(s string, subst string) (map[string]string, error) {
	var err error = nil
	optmap := make(map[string]string)
	for _, opt := range strings.Split(s, ",") {
		if idx := strings.Index(opt, "="); idx < 0 {
			optmap[opt] = subst
		} else {
			// Assing key as start-to-idx, skip '=', then value as idx-to-end
			optmap[opt[:idx]] = opt[idx+1:]
		}
	}

	return optmap, err
}

// ParseFlags parses flags of type below
//   -cpu model=ARMA50,freq=100,cache=256k
// first split on ',' then split on '='
//
func ParseFlags(s string) (map[string]string, error) {
	return _parseflags(s, "")
}

// ParseFlags parses flags of type below
//   -cpu model=ARMA50,freq=100,cache=256k
// first split on ',' then split on '='
// but case like 
//   -device pl011,iobase=0x210934
// returns first instance as map[subst]=pl011 
func ParseFlagsSubst(s string, subst string) (map[string]string, error) {
	return _parseflags(s, subst)
}


func ParseMem(v string) (uint64, error) {
	var mult, mem uint64
	var e error

	switch v[len(v)-1] {
	case 'k', 'K':
		mult = 1 << 10 // 2^10
	case 'm', 'M':
		mult = 1 << 20 // 2^20
	case 'g', 'G':
		mult = 1 << 30 // 2^30
	default:
		mult = 1
	}

	if mult != 0 {
		v = v[:len(v)-1]
	}

	if mem, e = strconv.ParseUint(v, 0, 64); e != nil {
		return 0, e
	}

	return mem * mult, nil
}

func ParseFreq(s string) (uint32, error) {
	var mult, mem uint64
	var e error

	if strings.HasSuffix(s, "hz") || strings.HasSuffix(s, "HZ") {
		s = s[:len(s)-2]
	}

	// See what the last char is
	switch s[len(s)-1] {
	case 'k', 'K':
		mult = 1e3 // 10^3
	case 'm', 'M':
		mult = 1e6 // 10^6
	case 'g', 'G':
		mult = 1e9 // 10^9
	default:
		mult = 1
	}

	if mult != 1 {
		s = s[:len(s)-1]
	}
	// Parse as Uint32 as processor freq are not more than 4GHz
	if mem, e = strconv.ParseUint(s, 0, 32); e != nil {
		return 0, e
	}

	return uint32(mem * mult), nil
}
