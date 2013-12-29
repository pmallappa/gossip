// Telnet Client program. Sample implementation
// Connects to telnet server,
// Reads user input and sends across network
// Writes network reads to STDOUT or io.Writer

package main

import (
	"fmt"
	"io"
	"os"
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

	err := tc.ConnectTimeout(proto, server, 200)
	if err != nil {
		fmt.Println(err)
		panic("Holla")
	}

	buf := make([]byte, 100)

	go io.Copy(os.Stdout, tc)

	for {
		n, err := os.Stdin.Read(buf)
		if err != nil {
			fmt.Printf("%s %s\n", buf[:n], err)
			break
		}

		n, err = tc.Write(buf[:n])
		if err != nil {
			fmt.Printf("Write Error: %s", err)
			break
		}

		time.After(2500 * time.Second)
	}

}
