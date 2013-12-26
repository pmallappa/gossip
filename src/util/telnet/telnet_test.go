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
	_, err := net.Dial(server, proto)
	//fmt.Printf("Client returning conn:%v err:%v\n", conn, err)
	c <- err
}

func Test_TelentConnection(t *testing.T) {
	telserve := NewServer()
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

func Test_ConnectionRefused(t *testing.T) {
	//telserve := NewTelnetServer()

	// Need to change address, as the old port may
	// not be available for sometime
	ch1 := make(chan error)
	server := ":2011"
	go dialAfter(ch1, proto, server, 7)
	//if e := telserve.ListenTimeoutProgress(proto, server, 3); e != nil {
	// We should see some error, verify its timeout error
	//}
	//telserve.Close()
	err := <-ch1
	if err == nil {
		t.Errorf("Something running on %s Client returned: %v", server, err)
	}
}

func dialAfter3(c chan error, server, proto, text string, t time.Duration) {
	time.After(t * time.Second)
	conn, err := net.Dial(server, proto)
	conn.Write([]byte(text))
	c <- err
}

func Test_Read(t *testing.T) {
	telserve := NewTelnetServer()
	server = ":2010"
	text := "Hello world\n"
	go dialAfter3(ch, proto, server, text, 1)
	if e := telserve.ListenTimeout(proto, server, 3); e != nil {
		t.Errorf("Error starting server %s", e)
	}
	line, _, _ := telserve.ReadLine()
	if string(line) != text[:len(text)-1] {
		fmt.Printf("recieved somethging else %s\n", line)
	}
	err := <-ch
	if err != nil {
		t.Errorf("Client returned: %v", err)
	}

}
