package telnet

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net"
	//"strings"
)

type Conn struct {
	net.Conn
	br *bufio.Reader
	bw *bufio.Writer
}

func newConn(conn net.Conn) (*Conn, error) {
	c := Conn{
		Conn: conn,
		bw:   bufio.NewWriterSize(conn, 512),
		br:   bufio.NewReaderSize(conn, 512),
	}
	return &c, nil
}

func Start(s []string) (io.Writer, error) {
	//
	// TODO: Need to fix this,
	// Required is, open a telnet port, keep listening
	// if a connection arrives, write the buffered output, and keep writing then on
	// if not, but a Write() requested, overwrite the old buffer
	//
	//recievechan := make(chan net.Conn)
	var (
		c       *Conn
		e       error
		netconn net.Conn
	)

	var listn net.Listener

	listn, e = net.Listen(s[0], s[1])
	if e != nil {
		return nil, e
	}

	//go func() {
	log.Println("Starting telnet server on: ", s[1])

	//	for {
	netconn, e = listn.Accept()
	if e != nil {
		log.Fatal("NOTHING")
	}
	if c, e = newConn(netconn); e != nil {
		return nil, e
	}

	//		recievechan <- c.conn
	//	}
	//}()
	//<-recievechan
	return c, nil
}

// io.Writer interface
func (c *Conn) Write(buf []byte) (int, error) {
	var pattern = []byte{'\n'}
	var (
		b, n int
		err  error
	)

	for len(buf) > 0 {
		if n, err = c.Conn.Write(buf); err != nil {
			return n, err
		}
		buf = buf[n:]
	}
	return n, err

	// =========== NEVER REACHED =================
	// Need to revisit,
	// With this we dont see data immediately, some debug cases
	// use UART to write debug data its not acceptible to loos
	// what is written on to console

	// Write using 'bw' buffered writer, on encountering '\n', need to do bw.Flush()

	for len(buf) > 0 {
		i := bytes.LastIndex(buf, pattern)
		if i < 0 {
			if n, err = c.bw.Write(buf); err != nil {
				return n, err
			}
			break
		}
		if b, err = c.bw.Write(buf[:i]); err != nil {
			c.bw.Write([]byte{'\n'})
			c.bw.Flush()
		}
		println("========>", buf)
		n += b
		buf = buf[b+1:]

		if b, err = c.bw.Write(buf); err != nil {
			return n + b + 1, err
		}
		println("========>", buf)
		n += b
		buf = buf[b:]
	}
	return n, err
}

// io.Reader interface
func (c *Conn) Read(buf []byte) (int, error) {
	var n int

	buflen := len(buf)
	for n < buflen {
		b, err := c.Conn.Read(buf)
		if err != nil {
			return b, err
		}
		//log.Printf("char: %d %s", b, buf)
		//buf[n] = b
		n += b
		buf = buf[b:]
		//if c.rwc.Buffered() == 0 {
		// Try don't block if can return some data
		//	break
		//}
	}
	return n, nil
}
