package main

import (
	"flag"
	"util/cflag"
)

func main() {
	myflag := cflag.NewStrFlag("flag")

	for _, v := range []*cflag.KVFlag{
		cflag.NewKVFlag("abcd", "Mydescription", 1),
		cflag.NewKVFlag("mycd", "MYnewdesc", "mystr"),
	} {
		myflag.Add(v)
	}

	flag.Var(myflag, "flag", "Hello how are you")
	flag.Parse()

}
