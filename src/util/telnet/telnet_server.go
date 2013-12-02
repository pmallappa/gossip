package telnet

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

type Server struct {
	Telnet
	Addr  string
	listn net.Listener
}

var defaultServer = &Server{}

func NewServer() *Server {
	return &Server{Addr: ":telnet"}
}

func connect(c chan error, t *Server) {
	var e error
	if t.conn, e = t.listn.Accept(); e != nil {
		c <- e
		return
	}

	t.bufwr = bufio.NewWriterSize(t.conn, 512)
	t.bufrd = bufio.NewReaderSize(t.conn, 512)

	c <- nil
}

func ListenAndServe() {
}

// Options are passed like telnet=tcp!localhost:2030
// Change is to accept everything that golang/pkg/net can do with
// 'proto' and 'addr'
// eg:
//        Dial("tcp", "12.34.56.78:80")      OR  Dial("tcp", "google.com:http")
//        Dial("tcp", "[2001:db8::1]:http")  OR  Dial("tcp", "[fe80::1%lo0]:80")
//        Dial("ip4:1", "127.0.0.1")         OR  Dial("ip6:ospf", "::1")
func (t *Server) ListenTimeout(proto, addr string, dur time.Duration) (e error) {
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
func (ts *Server) ListenTimeoutProgress(proto, addr string, dur time.Duration) (e error) {
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

func (ts *Server) Close() {
	if ts.debug {
		fmt.Printf("Closing: %v\n", ts.listn.Addr)
	}
	ts.Telnet.Close()
	ts.listn.Close()
}
