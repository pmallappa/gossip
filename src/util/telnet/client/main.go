// Telnet Client program. Sample implementation
// Connects to telnet server,
// Reads user input and sends across network
// Writes network reads to STDOUT or io.Writer

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

	tc := telnet.NewClient(proto, server)
	defer tc.Close()
	tc.EnableDebug()

	if err := tc.ConnectTimeout(proto, server, 200); err != nil {
		fmt.Println(err)
		panic("Holla")
	}

	// for {
	{ // We run only once
		tc.Write([]byte("ls /"))
		line, err := tc.ReadBytes(0)
		if err != nil {
			break
		}
		fmt.Print(string(line))
	}
}
