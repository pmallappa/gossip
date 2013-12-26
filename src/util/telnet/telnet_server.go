package telnet

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type serverT struct {
	telnetT
	proto, laddr string
	exec         string // Program to be invoked, after successful connection
	listn        net.Listener
}

type Server interface {
	Listen(proto, addr string) // A one time listener
	ListenAndServe(proto, addr string)
	ListenAndTimeout(proto, addr string, dur time.Duration) // Wait for connection till timeout
}

var defaultServer = &serverT{
	proto: "tcp",
	laddr: ":telnet",
	exec:  "/bin/sh",
}

func NewServerDefault() *serverT {
	return &defaultServer
}

func NewServer(proto, laddr string) *serverT {
	return &serverT{
		proto: proto,
		laddr: laddr,
	}
}

func connect(c chan error, t *serverT) {
	var e error
	if t.conn, e = t.listn.Accept(); e != nil {
		c <- e
		return
	}

	t.bufwr = bufio.NewWriterSize(t.conn, 512)
	t.bufrd = bufio.NewReaderSize(t.conn, 512)

	c <- nil
}

func handleConnection(exec string, conn *net.Conn) {
	// TODO:
	// Start a program denoted by 'exec',
	// untill the program exits or connection closes
	//    -> read connection, write to program input
	//    -> read program output, write to connection
	for {
		io.Copy(c, c)
	}
}

// This is a continous listener
func (t *serverT) ListenAndServe(proto, addr string) {

	if proto == "" {
		proto = t.proto
	}
	if addr == "" {
		addr = t.addr
	}

	ln, err := net.Listen(proto, addr)
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(t.exec, conn)
	}
}

// Options are passed like telnet=tcp!localhost:2030
// Change is to accept everything that golang/pkg/net can do with
// 'proto' and 'addr'
// net.Listen("tcp", ":8080")

func (t *serverT) _listenTimeout(proto, addr string, dur time.Duration) (e error) {
	if t.listn, e = net.Listen(proto, addr); e != nil {
		return
	}

	con_ch := make(chan error)
	go connect(con_ch, t)

	select {
	case <-time.After(dur * time.Second):
	case e = <-con_ch:
	}

	if e != nil {
		fmt.Printf("%v", e)
	}

	if t.debug {
		fmt.Printf("\n%v", addr)
		if t.conn != nil {
			fmt.Println("[Connected]")
		}
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
		if ts.conn != nil { // We got a connection
			return nil
		}
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
	ts.telnetT.Close()
	ts.listn.Close()
}
