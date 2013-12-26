package telnet

import (
	"net"
	"time"
)

// proto and addr are as follows
// eg:
//        Dial("tcp", "12.34.56.78:80")      OR  Dial("tcp", "google.com:http")
//        Dial("tcp", "[2001:db8::1]:http")  OR  Dial("tcp", "[fe80::1%lo0]:80")
//        Dial("ip4:1", "127.0.0.1")         OR  Dial("ip6:ospf", "::1")

type clientT struct {
	telnetT
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
	Connect(proto, addr string) *net.Conn
	ConnectTimeout(proto, addr string, dur time.Duration) *net.Conn
}
