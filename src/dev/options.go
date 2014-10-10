package dev

import (
	"flag"
	//"fmt"
	//"strings"
)

import (
	//	"util"
	"util/cflag"
)

/*
-drive if=none,id=DRIVE-ID,HOST-OPTS...
-device DEVNAME,drive=DRIVE-ID,DEV-OPTS...
*/

type Devopt []*cflag.FlagSet

func (d *Devopt) Set(s string) error {
	o := cflag.NewFlagSet(s, flag.ExitOnError)
	*d = append(*d, o)
	return nil
}

func (d *Devopt) Get() string {
	return ""
}

func (d *Devopt) String() string {
	return ""
}

var (
	deviceOpts,
	driveOpts, // Used for block devices
	chrdevOpts, // Any character device
	netdevOpts Devopt // Netdevice
)

func Parse(str string) (map[string]string, error) {
	var e error
	var m map[string]string

	return m, e
}

func initDevFlagOptions() {
	/*
		for _, v := range []opt{
			{"drive", "", "Uniquely identify a device node"},
			{"chardev", "", "Identify a chardev"},
			{"netdev", "", "Identify a netdev"},
		} {
			deviceOpts.c.Add(cflag.NewSubOption(v.name, v.desc, v.name))
		}
	*/
}

func initDevFlags() {
	flag.Var(&deviceOpts, "device", "Device part, of a Block device")

	flag.Var(&driveOpts, "drive", "Drive part, of a block device")
	flag.Var(&chrdevOpts, "chardev", "Host part of options for a chardev")
	flag.Var(&netdevOpts, "netdev", "Host part of netdev options")
}
