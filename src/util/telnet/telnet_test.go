package telnet

import (
	"testing"
)

import (
	"net"
	"time"
)

func dialafter1sec(server string, proto string) (err error) {
	time.After(1 * time.Second)
	_, err = net.Dial(server, proto)
	if err != nil {
		return err
	}
	return err
}

func Test_Telent(t *testing.T) {
	tel := NewTelnet()
	server := ":2000"
	proto := "tcp"
	go dialafter1sec(proto, server)
	if e := tel.ConnectTimeout(proto, server, 3); e != nil {
		t.Errorf("Halla bol")
	}

}
