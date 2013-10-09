package telnet

import (
	"testing"
)

import (
	//"fmt"
	"net"
	"time"
)

var (
	ch     = make(chan error)
	server = ":2000"
	proto  = "tcp"
)

func dialAfter(c chan error, server string, proto string, t time.Duration) {
	var err error
	time.After(t * time.Second)
	_, err = net.Dial(server, proto)
	c <- err
}

func Test_Telent(t *testing.T) {
	tel := NewTelnet()
	defer tel.Close()
	server := ":2000"
	go dialAfter(ch, proto, server, 1)
	if e := tel.ListenTimeout(proto, server, 3); e != nil {
		t.Errorf("Error starting server %s", e)
	}
	e := <-ch
	if e != nil {
		t.Errorf("Client returned: %v", e)
	}

}

func Test_Telnet1(t *testing.T) {
	tel := NewTelnet()
	defer tel.Close()
	// Need to change address, as the old port may
	// not be available for sometime
	server := ":2001"
	go dialAfter(ch, proto, server, 4)
	if e := tel.ListenTimeoutProgress(proto, server, 3); e != nil {
		// We should see some error, verify its timeout error
	}
	e := <-ch
	if e == nil {
		t.Errorf("Client returned: %v", e)
	}
}
