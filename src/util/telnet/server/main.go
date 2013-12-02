package main

import (
	"fmt"
	//"time"
)
import (
	"util/telnet"
)

var (
	//ch     = make(chan error)
	server = ":2000"
	proto  = "tcp"
)

func main() {

	ts := telnet.NewTelnetServer()
	defer ts.Close()
	ts.EnableDebug()

	if err := ts.ListenTimeoutProgress(proto, server, 20); err != nil {
		fmt.Println(err)
		panic("Holla")
	}

	for {
		line, err := ts.ReadBytes(0)
		if err != nil {
			break
		}
		fmt.Print(string(line))
	}
}
