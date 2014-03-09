package plat

import (
	"errors"
	"flag"
	"fmt"
	//"strconv"
	//"strings"
)

import (
//	"util"
)

type loadReqs struct {
	addr  uint64
	value interface{}
}

type loadOptstr []string

func (f *loadOptstr) String() string {
	return fmt.Sprint([]string(*f))
}

func (f *loadOptstr) Set(value string) error {
	*f = append(*f, value)
	return nil
}

//
// -ld type=file,addr=file=<file>    - Loads raw contents of file to addr
// -ld type=elf,file=<file>          - Loads an ELF file, address is as per ELF sections
// -ld type=zero,addr=,len=          - Loads zero's at addr
//

var (
	ldOpts    loadOptstr
	load_opts loadOptstr

	ld_help_str string = "type=file,file=<file>,addr=    - Loads raw contents of file to addr\n" +
		"\t type=elf,file=<file>    - Loads an ELF file, address is as per ELF sections\n" +
		"\t type=zero,addr=,len=    - Loads zero's at 'addr'\n" +
		"\t Conflicting(duplicate) address will not be detected, but can be figured out" + "from Loader prints"

	elfloads  []string
	zeroloads map[uint64]uint64 //addr : size
	fileloads map[uint64]string //addr : file
)

func init() {
	fileloads = make(map[uint64]string)
	zeroloads = make(map[uint64]uint64)
	elfloads = make([]string, 0, 128)

	flag.Var(&load_opts, "ld", ld_help_str)
}

func loadUsage() {
	println(ld_help_str)
}

func parseLoadFlags() error {

	for _, entry := range load_opts {
		var m map[string]string
		//var e error
		//if m, e = util.ParseFlagsSubst(entry, "type"); e != nil {
		//	return errors.New("Unable to Parse")
		//}
		switch m["type"] {
		case "file":
			if _, present := m["name"]; !present {
				if _, present := m["addr"]; present {
				} else {
					return errors.New("Should provide addr=<address> for file")
				}
			} else {
				return errors.New("Should provide name=<name> for file")
			}

			//if addr, err := util.ParseMem(m["addr"]); err == nil {
			//	fileloads[addr] = m["name"]
			//} else {
			//	return err
			//}
		case "elf":
			elfloads = append(elfloads, entry)
		case "zero":
			if _, present := m["len"]; !present {
				if _, present := m["addr"]; !present {

				} else {
					return errors.New("Should provide len=<length in Bytes> K,M,G suffix is supported")
				}
			} else {
				return errors.New("Should provide addr=<addr> K, M, G suffix supported")
			}

			//if length, err := util.ParseMem(m["len"]); err == nil {
			//	if addr, err := util.ParseMem(m["addr"]); err == nil {
			//		zeroloads[addr] = length
			//	} else {
			//		return err
			//	}
			//} else {
			//	return errors.New("Should provide addr=<address> for file")
			//}

		default:
			return fmt.Errorf("Dont understand -ld option %s", m["type"])
		}
	}
	return nil
}

func initLdOpts() {
	flag.Var(&ldOpts, "Load options", "")
}
