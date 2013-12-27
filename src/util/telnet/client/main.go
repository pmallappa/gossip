// Telnet Client program. Sample implementation
// Connects to telnet server,
// Reads user input and sends across network
// Writes network reads to STDOUT or io.Writer

package main

import (
	"fmt"
	//"io"
	"time"
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
	buf := make([]byte, 100)
	for {
		println("clent: Before write")
		tc.Write([]byte("/usr/bin/fdisk -l"))
		for {
			println("client: Before read")
			n, err := tc.ReadLine(buf)
			if err != nil || n == 0 {
				break
			}
			fmt.Printf("%s\n", buf[:n])
		}
		time.After(2*time.Second)
	}
}
