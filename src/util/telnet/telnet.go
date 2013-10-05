package telnet

import (
	//	"bytes"
	"errors"
	"fmt"
	"net"
	"os"
	//	"strconv"
	"time"
)

type Telnet struct {
	in, out         chan []byte
	done            chan bool
	buf             []byte
	Conn            net.Conn
	ignConnectError bool
}

func (t *Telnet) IgnoreErrors() {
	t.ignConnectError = true
}

func (t *Telnet) Read(buf []byte) (int, error) {
	if t.Conn == nil { // We might have dummy connection ignore reads
		return 0, nil
	}
	n, err := t.Conn.Read(buf)
	if err != nil {
		return n, err
	}
	fmt.Println("Read", string(buf))
	return len(buf), nil
}

func (t *Telnet) Write(buf []byte) (int, error) {
	if t.Conn == nil {
		return 0, nil
	}
	n, err := t.Conn.Write(buf)
	if err != nil {
		fmt.Printf("Write:Error %v", err)
	}
	return n, nil
}

func NewTelnet() *Telnet {
	t := Telnet{
		in:   make(chan []byte), // TODO: make it recieve only
		out:  make(chan []byte), // TODO: make it send only
		done: make(chan bool),
		buf:  make([]byte, 64),
	}
	return &t
}

// wait for incoming connection only first time,
// After connection established and closed, All writes are ignored
// All reads return with bytes read as 0
func connect(c chan bool, listener net.Listener, t *Telnet) {
	var e error
	t.Conn, e = listener.Accept()
	if e != nil {
		fmt.Printf("connect: Error: %v", e)
		os.Exit(128)
	}
	c <- true
}

func (t *Telnet) ConnectTimeout(proto string, addr string,
	timeout uint32) error {
	listener, e := net.Listen(proto, addr)
	if e != nil {
		fmt.Printf("Telnet Listen Error %v", e)
		return e
	}

	con := make(chan bool)

	go connect(con, listener, t)

	for ; timeout > 0; timeout-- {
		var connected bool
		fmt.Printf("Waiting %d seconds for connection \r", timeout)
		select {
		case <-time.After(1 * time.Second):

		case connected = <-con:
		}
		if connected {
			fmt.Printf("\n")
			break
		}
	}

	if timeout == 0 {
		e = errors.New("Connection timed out\n")
	}
	if t.ignConnectError {
		e = nil
	}
	return e
}

func (t *Telnet) Connect(proto string, addr string) error {
	return t.ConnectTimeout(proto, addr, 0)
}

func (t *Telnet) Start() error {
	return nil
}
