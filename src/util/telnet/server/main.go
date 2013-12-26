// Telnet Server program, Sample Implementation
// Serves at given port, or configured port
// Listens, upon connection spawns off a connection handler
//	-> The default connection handler 'exec's a shell
//		-> Connects shell's output to connection(io.Writer)
//	-> Reads lines from input(with enabled telnet protocol control chars)
//		-> repeats character reading untill line is read
//	-> All the input is given to shell,
//		-> and shell output is written back to connection

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

	ts := telnet.NewServer()
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
