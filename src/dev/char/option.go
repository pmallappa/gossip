package char

import (
	"util/cflag"
)

var chardevopts = cflag.NewCFlagSet("chardev")

func initCharOpts() error {

	for _, v := range []*cflag.Suboption{
		cflag.NewSubOptionOther(&new(int), "ioaddr", "IO Address", "0"),
		cflag.NewSubOptionOther(&new(int), "irq", "IRQ", "0"),
		cflag.NewSubOptionOther(&new(int), "ioaddr", "IO Address", "0"),
		cflag.NewSubOptionString(&new(string), "name", "Name", ""),
		cflag.NewSubOptionString(&new(int), "id", "Identifier", "0"),
	} {
		chardevopts.Add(v)
	}

	flag.Var(chardevopts, "chardev", "Charcter Device options")
	return nil
}
