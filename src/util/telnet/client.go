package telnet

import (
	"bufio"
	"net"
	"time"
)

// proto and addr are as follows
// eg:
// Dial("tcp", "12.34.56.78:80")      OR  Dial("tcp", "google.com:http")
// Dial("tcp", "[2001:db8::1]:http")  OR  Dial("tcp", "[fe80::1%lo0]:80")
// Dial("ip4:1", "127.0.0.1")         OR  Dial("ip6:ospf", "::1")

type clientT struct {
	telnetT
	bufrd       *bufio.Reader
	proto, addr string
}

var defaultClient = clientT{
	proto: "tcp",
	addr:  ":telnet",
}

func NewClientDefault() *clientT {
	return &defaultClient
}

func NewClient(proto, laddr string) *clientT {
	return &clientT{
		proto: proto,
		addr:  laddr,
	}
}

type Client interface {
	Connect(proto, addr string) error
	ConnectTimeout(proto, addr string, dur time.Duration) error
}

func (tc *clientT) Connect(proto, addr string) (e error) {
	if tc.conn, e = net.Dial(proto, addr); e != nil {
		return
	}
	//tc.bufwr = bufio.NewWriterSize(tc.conn, 512)
	tc.bufrd = bufio.NewReaderSize(tc.conn, 512)

	return
}

func (tc *clientT) ConnectTimeout(proto, addr string, t time.Duration) (e error) {
	return tc.Connect(proto, addr)
}

// TODO, this has to go through telnetT.Read()
func (tc *clientT) Read(buf []byte) (n int, e error) {
	n, e = tc.bufrd.Read(buf)
	return
}
