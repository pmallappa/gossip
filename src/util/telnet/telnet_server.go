package telnet

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os/exec"
	"strings"
	"time"
)

type serverT struct {
	//telnetT
	proto, laddr string
	exec         string // after successful connection
	debug        bool
	listn        net.Listener
}

func (ts *serverT) EnableDebug() {
	ts.debug = true
}

type Server interface {
	Listen(proto, addr string) // A one time listener
	ListenAndServe(proto, addr string)

	// Wait for connection till timeout
	ListenTimeout(proto, addr string, dur time.Duration)
}

var defaultServer = serverT{
	proto: "tcp",
	laddr: ":telnet",
	exec:  "/usr/bin/sh",
}

func NewServerDefault() *serverT {
	return &defaultServer
}

func NewServer(proto, laddr string) *serverT {
	return &serverT{
		proto: proto,
		laddr: laddr,
		exec:  "/usr/bin/sh",
	}
}

// Start a program denoted by 'exec',
// untill the program exits or connection closes
//    -> read connection, write to program input
//    -> read program output, write to connection
func handleConnection(t *telnetT, command string) {

	var err error

	defer t.Close()

	t.bufwr = bufio.NewWriterSize(t.conn, 512)
	t.bufrd = bufio.NewReaderSize(t.conn, 512)

	split := strings.Split(command, " ")
	cmd := exec.Command(split[0])
	if len(split) > 1 {
		cmd.Args = split[1:]
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	if t.debug {
		fmt.Printf("Starting command :%s", command)
	}

	if err = cmd.Start(); err != nil {
		return
	}

	go io.Copy(stdin, t.bufrd)
	go io.Copy(t.bufwr, stdout)

	err = cmd.Wait()
	if t.debug {
		fmt.Printf("server: %s exited with %v", command, err)
	}

	fmt.Printf("%s\n", err)
}

// This is a continous listener
func (st *serverT) ListenAndServe(proto, addr string) (err error) {

	if proto == "" {
		proto = st.proto
	}
	if addr == "" {
		addr = st.laddr
	}

	ln, err := net.Listen(proto, addr)
	if err != nil {
		return err
	}
	for {
		t := NewTelnet()
		if st.debug {
			t.EnableDebug()
		}
		t.conn, err = ln.Accept()
		if err != nil {
			// handle error
			continue
		}

		go handleConnection(t, st.exec)
	}
	return nil
}

// Options are passed like telnet=tcp!localhost:2030
// Change is to accept everything that golang/pkg/net can do with
// 'proto' and 'addr'
// net.Listen("tcp", ":8080")

func (t *serverT) ListenTimeout(proto, addr string, dur time.Duration) (e error) {
	if t.listn, e = net.Listen(proto, addr); e != nil {
		return
	}

	con_ch := make(chan error)
	//go connect(con_ch, t)

	select {
	case <-time.After(dur * time.Second):
	case e = <-con_ch:
	}

	if e != nil {
		fmt.Printf("%v", e)
	}

	if t.debug {
		fmt.Printf("\n%v", addr)
		// 	if t.conn != nil {
		// 		fmt.Println("[Connected]")
		// 	}
	}
	return
}

// This is similar to ListenTimeout, but it prints number of seconds waited,
// And exits if Errors are supposed to be treated seriously
func (ts *serverT) _listenTimeoutProgress(proto, addr string, dur time.Duration) (e error) {
	timeout := int(dur)
	//c := make(chan bool)
	//go ts._counter(c, timeout)

	go ts.ListenTimeout(proto, addr, dur)

	for ; timeout > 0; timeout-- {
		fmt.Printf("Waiting %d seconds for connection \r", timeout)
		<-time.After(1 * time.Second)
		//if ts.conn != nil { // We got a connection
		//	return nil
		//}
	}

	if timeout == 0 {
		return fmt.Errorf("Timed out")
	}
	return
}

func (ts *serverT) Close() {
	if ts.debug {
		fmt.Printf("Closing: %v\n", ts.listn.Addr)
	}
	ts.listn.Close()
}
