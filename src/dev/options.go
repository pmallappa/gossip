package dev

import (
	"flag"
	"fmt"
	"strings"
)

import (
	//	"util"
	"util/cflag"
)

type devopt string

var (
	devOpts     devopt
	charDevOpts = cflag.New()
	netDevOpts  = cflag.New()
	usbDevOpts  = cflag.New()
)

func Parse(str string) (map[string]string, error) {
	var e error
	var m map[string]string

	return m, e
}

func initDevFlags() {
	flag.Var(&devOpts, "dev", "Generic device help")
	flag.Var(charDevOpts, "chardev", "Charcter devices, use ? for more")
	flag.Var(netDevOpts, "netdev", "Network Devices, use ? for more")
	flag.Var(usbDevOpts, "usbdev", "USB Devices, use ? for more")
}

func (f *devopt) String() string {
	return fmt.Sprint(string(*f))
}

func (f *devopt) Set(str string) error {
	// skip the prefix and set appropriate device
	switch {
	case strings.HasPrefix(str, "char,") == true:
		charDevOpts.Set(str[4:])
	case strings.HasPrefix(str, "net,") == true:
		netDevOpts.Set(str[3:])
	case strings.HasPrefix(str, "usb,") == true:
		usbDevOpts.Set(str[3:])
	default:
		fmt.Printf("Every device has to be prefix with char/net/usb")
	}
	return nil
}
