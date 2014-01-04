package dev

import (
	"flag"
	"fmt"
)

import (
	"util"
	"util/cflag"
)

type devOptstr []string

var (
	devOpts, chardevOpts,
	netdevOpts, pcidevOpts, usbdevOpts devOptstr
	devOpts_help     string = "-dev ?"
	chardevOpts_help string = "-chardev ?"
	netdevOpts_help  string = "-netdev ?"
	pcidevOpts_help  string = "-pcidev ?"
	usbdevOpts_help  string = "-usbdev ?"
)

func Parse(str string) (map[string]string, error) {
	var e error
	var m map[string]string

	for _, str := range devOpts {
		if m, e := util.ParseFlags(str); e != nil {
			return m, e
		}

		for k, v := range m {
			switch k {
			case "char":
				chardevOpts.Set(str)
			case "net":
				netdevOpts.Set(str)
			case "usb":
				usbdevOpts.Set(str)
			default:
				v = v
			}
		}
	}
	return m, e
}

func initDevFlags() {
	devOpts = make([]string, 0, 128)
	chardevOpts = make([]string, 0, 128)
	netdevOpts = make([]string, 0, 128)
	pcidevOpts = make([]string, 0, 128)
	usbdevOpts = make([]string, 0, 128)

	flag.Var(&devOpts, "dev", devOpts_help)
	flag.Var(&chardevOpts, "chardev", chardevOpts_help)
	flag.Var(&netdevOpts, "netdev", netdevOpts_help)
	flag.Var(&pcidevOpts, "pcidev", pcidevOpts_help)
	flag.Var(&usbdevOpts, "usbdev", usbdevOpts_help)
}

func (f *devOptstr) String() string {
	return fmt.Sprint([]string(*f))
}

func (f *devOptstr) Set(value string) error {
	*f = append(*f, value)
	return nil
}
