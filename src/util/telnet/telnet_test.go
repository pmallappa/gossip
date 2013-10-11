package telnet

import (
	"testing"
)

import (
	"fmt"
	"net"
	"time"
)

var (
	ch     = make(chan error)
	server = ":2000"
	proto  = "tcp"
)

func dialAfter(c chan error, server string, proto string, t time.Duration) {
	time.After(t * time.Second)
	conn, err := net.Dial(server, proto)
	fmt.Printf("Client returning conn:%v err:%v\n", conn, err)
	c <- err
}

func Test_Telent(t *testing.T) {
	telserve := NewTelnetServer()
	defer telserve.Close()
	server := ":2000"
	go dialAfter(ch, proto, server, 1)
	if e := telserve.ListenTimeout(proto, server, 3); e != nil {
		t.Errorf("Error starting server %s", e)
	}
	err := <-ch
	if err != nil {
		t.Errorf("Client returned: %v", err)
	}

}

func Test_Telnet1(t *testing.T) {
	//telserve := NewTelnetServer()

	// Need to change address, as the old port may
	// not be available for sometime
	ch1 := make(chan error)
	server := ":2001"
	go dialAfter(ch1, proto, server, 7)
	//if e := telserve.ListenTimeoutProgress(proto, server, 3); e != nil {
	// We should see some error, verify its timeout error
	//}
	//telserve.Close()
	err := <-ch1
	if err == nil {
		t.Errorf("Client returned: %v", err)
	}

}

// func Test_Telnet2(t *testing.T) {
// 	server = ":3000"
// 	go dialAfter(ch, proto, server, 1)
// 	_ = <-ch
// }
