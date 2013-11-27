// Package telnet,
// Implements RFC 854
// Features:
//     - Implements client related functions
//     - Server related
//     - Generic io.Reader and io.Writer to connect to any other programs

//  RFC 854  Telnet Protocol Specification
//  RFC 855  Telnet Option Specifications
//  RFC 856  Telnet Binary Transmission
//  RFC 857  Telnet Echo Option
//  RFC 858  Telnet Suppress Go Ahead option
//  RFC 859  Telnet Status Option
//  RFC 860  Telnet Timing Mark option
//  RFC 861  Telnet Extended Options option
//  RFC 1091 Telnet Terminal Type option
//  RFC 1096 Telnet X Display location option
//  RFC 1184 Telnet Linemode option
//  RFC 1372 Telnet Remote Flow option
//  RFC 1408 Telnet Environment option
//  RFC 1571 Telnet Environment option interoperability issues
//  RFC 1572 Telnet Environment option
//  RFC 2066 Telnet Charset option
//  RFC 2941 Telnet Authentication option
//  RFC 2840 Telnet Kermit option
//  RFC 2217 Telnet Com Port option
//  RFC 1073 Telnet Window Size option
//  RFC 1079 Telnet Terminal Speed option
//  RFC 727  Telnet logout option

package telnet

import (
	"bufio"
	//"bytes"
	"fmt"
	//"io"
	//"log"
	"net"
	//"strings"
	"time"
)

// Each sequence of read is one or two bytes, depending on if the byte read
// has value {0-254} or 255, if later(255) case, then its a control command,
// and next byte indicates the actual command

type Telnet struct {
	conn     net.Conn
	bufrd    *bufio.Reader
	bufwr    *bufio.Writer // Dont know if buffered writer required
	unixCRLF bool

	debug bool
}

func (t *Telnet) EnableDebug() {
	t.debug = true
}

type TelnetServer struct {
	Telnet
	listn net.Listener
}

const (
	CR  = byte('\r')
	LF  = byte('\n')
	NUL = byte(0)
)

type telnetCMD byte

const (
	cmd_EOF        telnetCMD = 236 + iota // End of file
	cmd_SUSP                              // 237: Suspend process
	cmd_ABORT                             // Abort process
	cmd_EOR                               // end of record
	cmd_NOP                               // 240: nop
	cmd_DM                                // data mark
	cmd_BREAK                             // break
	cmd_IP                                // interrupt process
	cmd_AO                                // abort output
	cmd_AYT                               // 245: Are You There
	cmd_EC                                // delete current character
	cmd_EL                                // delete current line
	cmd_GA                                // Line reverse
	cmd_SE                                // end of sub negotiation
	cmd_SB                                // 250: Subnegotiation
	cmd_WILL                              // Indicating Option *WILL* be used
	cmd_WONT                              // Indicating option *WONT* be used
	cmd_DO                                // Commanding, to use Option
	cmd_DONT                              // Response, Option not supported
	cmd_IAC                               // 255: Interpret As Command
	cmd_FIRSTENTRY = cmd_EOF
)

var cmdStrings = []string{
	"EOF", "SUSP", "ABORT", "EOR",
	"NOP", "DM", "BREAK", "IP",
	"AO", "AYT", "EC", "EL", "GA",
	"SE", "SB", "WILL", "WONT", "DO",
	"DONT", "IAC",
}

type telnetOPT byte

const (
	opt_BINARY         telnetOPT = iota // 8-bit data path
	opt_ECHO                            // echo
	opt_RCP                             // prepare to reconnect
	opt_SGA                             // suppress go ahead
	opt_NAMS                            // approximate message size
	opt_STATUS                          // give status
	opt_TM                              // timing mark
	opt_RCTE                            // remote controlled transmission and echo
	opt_NAOL                            // negotiate about output line width
	opt_NAOP                            // negotiate about output page size
	opt_NAOCRD                          // negotiate about CR disposition
	opt_NAOHTS                          // negotiate about horizontal tabstops
	opt_NAOHTD                          // negotiate about horizontal tab disposition
	opt_NAOFFD                          // negotiate about formfeed disposition
	opt_NAOVTS                          // negotiate about vertical tab stops
	opt_NAOVTD                          // negotiate about vertical tab disposition
	opt_NAOLFD                          // negotiate about output LF disposition
	opt_XASCII                          // extended ascic character set
	opt_LOGOUT                          // force logout
	opt_BM                              // byte macro
	opt_DET                             // data entry terminal
	opt_SUPDUP                          // supdup protocol
	opt_SUPDUPOUTPUT                    // supdup output
	opt_SNDLOC                          // send location
	opt_TTYPE                           // terminal type
	opt_EOR                             // end or record
	opt_TUID                            // TACACS user identification
	opt_OUTMRK                          // output marking
	opt_TTYLOC                          // terminal location number
	opt_3270REGIME                      // 3270 regime
	opt_X3PAD                           // X.3 PAD
	opt_NAWS                            // window size
	opt_TSPEED                          // terminal speed
	opt_LFLOW                           // remote flow control
	opt_LINEMODE                        // Linemode option
	opt_XDISPLOC                        // X Display Location
	opt_OLD_ENVIRON                     // Old - Environment variables
	opt_AUTHENTICATION                  // Authenticate
	opt_ENCRYPT                         // Encryption option
	opt_NEW_ENVIRON                     // New - Environment variables
	opt_EXOPL          = 255            // extended-options-list
)

var optStrings = []string{
	"BINARY", "ECHO", "RCP", "SGA", "NAMS",
	"STATUS", "TM", "RCTE", "NAOL", "NAOP",
	"NAOCRD", "NAOHTS", "NAOHTD", "NAOFFD",
	"NAOVTS", "NAOVTD", "NAOLFD", "XASCII",
	"LOGOUT", "BM", "DET", "SUPDUP", "SUPDUPOUTPUT",
	"SNDLOC", "TTYPE", "EOR", "TUID", "OUTMRK",
	"TTYLOC", "3270REGIME", "X3PAD", "NAWS",
	"TSPEED", "LFLOW", "LINEMODE", "XDISPLOC",
	"OLD_ENVIRON", "AUTHENTICATION", "ENCRYPT",
	"NEW_ENVIRON", "EXOPL",
}

//
//
// ######    ##    #####      ##
//   ##    ##  ##  ##  ##   ##  ##
//   ##    ##  ##  ##   ##  ##  ##
//   ##    ##  ##  ##  ##   ##  ##
//   ##      ##    #####      ##
//  Some options have sub-options,
//     opt_AUTHENTICATION, opt_ENCRYPT,
//  And will be implemented in course of time.
//

func NewTelnet() *Telnet {
	return &Telnet{
		unixCRLF: true,
	}
}

func NewTelnetServer() *TelnetServer {
	return &TelnetServer{}
}

func connect(c chan error, t *TelnetServer) {
	var e error
	if t.conn, e = t.listn.Accept(); e != nil {
		c <- e
		return
	}

	t.bufwr = bufio.NewWriterSize(t.conn, 512)
	t.bufrd = bufio.NewReaderSize(t.conn, 512)

	c <- nil
}

// Options are passed like telnet=tcp!localhost:2030
// Change is to accept everything that golang/pkg/net can do with
// 'proto' and 'addr'
// eg:
//        Dial("tcp", "12.34.56.78:80")      OR  Dial("tcp", "google.com:http")
//        Dial("tcp", "[2001:db8::1]:http")  OR  Dial("tcp", "[fe80::1%lo0]:80")
//        Dial("ip4:1", "127.0.0.1")         OR  Dial("ip6:ospf", "::1")
func (t *TelnetServer) ListenTimeout(proto, addr string, dur time.Duration) (e error) {
	if t.listn, e = net.Listen(proto, addr); e != nil {
		return
	}

	con_ch := make(chan error)
	go connect(con_ch, t)

	select {
	case <-time.After(dur):
	case e = <-con_ch:
	}

	if e != nil {
		fmt.Printf("%v", e)
	}

	if t.debug {
		if t.conn != nil {
			fmt.Printf("Connected:")
		}
		fmt.Printf("listn: %v, conn:%v\n", t.listn, t.conn)
	}
	return
}

// This is similar to ListenTimeout, except, it prints number of seconds waited,
// And exits if Errors are supposed to be treated seriously
func (t *TelnetServer) ListenTimeoutProgress(proto, addr string, dur time.Duration) (e error) {
	timeout := int(dur)
	for ; timeout > 0; timeout-- {
		if t.debug {
			fmt.Printf("Waiting %d seconds for connection \n", timeout)
		}
		if e = t.ListenTimeout(proto, addr, 1); e != nil {
			return
		}
		// In case we are connected we are out
		if t.conn != nil {
			return
		}
	}

	return
}

// io.Writer interface
func (t *Telnet) Write(buf []byte) (int, error) {
	var (
		n   int
		err error
	)

	for len(buf) > 0 {
		if n, err = t.conn.Write(buf); err != nil {
			return n, err
		}
		buf = buf[n:]
	}
	return n, err
}

// io.Reader interface
func (t *Telnet) Read(buf []byte) (int, error) {
	var n int

	buflen := len(buf)
	for n < buflen {
		b, err := t.conn.Read(buf)
		if err != nil {
			return b, err
		}
		n += b
		buf = buf[b:]
	}
	return n, nil
}

func (t *Telnet) __execCMD(c telnetCMD) (err error) {
	//if t.debug {
	fmt.Printf("Command recieved %s\n", optStrings[c])
	//}
	switch c {
	case cmd_ABORT:
	case cmd_SUSP:
	case cmd_EOR:
	case cmd_NOP:
	case cmd_DM:
	case cmd_BREAK:
	case cmd_IP:
	case cmd_AO:
	case cmd_AYT:
	case cmd_EC:
	case cmd_EL:
	case cmd_GA:
	case cmd_SE:
	case cmd_SB:
	case cmd_WILL:
	case cmd_WONT:
	case cmd_DO:
	case cmd_DONT:
	}
	return
}

func (t *Telnet) __readByte() (c byte, err error) {
	if c, err = t.bufrd.ReadByte(); err != nil {
		return 0, err
	}

	cmd := telnetCMD(c)
	if cmd == cmd_IAC {
		err = t.__execCMD(cmd)
	}

	return
}

// bufio.Reader
func (t *Telnet) ReadByte() (c byte, err error) {
	// TODO: We have to interpret the 'telnet' commands and options
	// Send the left overs to whoever asking
	for {
		if c, err = t.__readByte(); err != nil {
			c = 0
			break
		}
	}

	return
}

func (t *Telnet) ReadBytes(delim byte) (line []byte, err error) {
	// TODO: need to take care of interpreting the commands
	if delim == 0 {
		delim = LF
	}

	for {
		c, err := t.ReadByte()
		if err != nil {
			break
		}
		line = append(line, c)
	}
	return
}

func (t *Telnet) ReadLine() (line []byte, isPrefix bool, err error) {
	return t.bufrd.ReadLine()
}

// func (t *Telnet) ReadRune() (r rune, size int, err error)           {}
// func (t *Telnet) ReadSlice(delim byte) (line []byte, err error)     {}
// func (t *Telnet) ReadString(delim byte) (line string, err error)    {}

func (t *Telnet) Close() {
	if t.debug {
		fmt.Printf("Closing conn:%v\n", t.conn)
	}
	t.conn.Close()
}

func (ts *TelnetServer) Close() {
	//if ts.debug {
	//	fmt.Printf("Closing listn:%v conn:%v\n", ts.listn, ts.conn)
	//}
	ts.Telnet.Close()
	ts.listn.Close()
}
