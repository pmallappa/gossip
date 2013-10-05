package logng

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

import (
	"util/telnet"
)

type LoggerType uint32

const (
	TcpUdp LoggerType = 1 << iota
	File
)

type LogLevel uint32

const (
	PANIC LogLevel = iota
	FATAL
	CRITICAL
	ERROR
	WARNING
	INFO
	DEBUG
)

var logLevelString = []string{
	PANIC:    "Panic",    // Unfavourable component error
	FATAL:    "Fatal",    // Internal simulator Error
	CRITICAL: "Critical", // Component, needing attention
	ERROR:    "Error",    // Component, reporting Error
	WARNING:  "Warning",  // Component reporting warning
	INFO:     "Info",     //
	DEBUG:    "Debug",    //
}

type LoggerNG struct {
	log.Logger
	str         string
	logType     LoggerType
	component   string        // Component
	fn          func() string // Dynamically Call a function for string
	curLevel    LogLevel
	exitOnError bool
}

func (log *LoggerNG) SetLevel(l LogLevel) {
	log.curLevel = l
}

func (log *LoggerNG) SetComponent(s string) {
	log.component = s
}

//
// ParseLogger understands any of the following options
//				log=file:<path>
//				log=file
//				log=tcp:address:port
// for the second option we rely on the 'util/telnet' module
func ParseLogger(s string) (*LoggerNG, error) {
	loggerstr := strings.SplitN(s, ":", 2)

	l := NewLoggerNG()
	l.logType = File

	switch strings.ToLower(loggerstr[0]) {
	case "tcp", "udp":
		l.logType = TcpUdp
	case "", "file":
	default:
		return nil, fmt.Errorf("Unknown description for logger '%s'", s)
	}

	l.str = loggerstr[1]

	return l, nil
}

func (l *LoggerNG) InitLogger() (*log.Logger, error) {
	var (
		logwriter io.Writer
		e         error
	)

	switch l.logType {
	case TcpUdp:
		t := telnet.NewTelnet()
		s := strings.SplitN(l.str, ":", 2)
		if s[1] == "" {
			s[1] = "21"
		}
		if s[0] == "" {
			s[0] = "localhost"
		}
		if e = t.Connect(s[0], s[1]); e != nil {
			return nil, e
		}
	case File:
		if logwriter, e = os.OpenFile(l.str, os.O_WRONLY|os.O_CREATE,
			0640); e != nil {
			return nil, e
		}
	default:
		logwriter = os.Stderr
	}

	return log.New(logwriter, "", 0), e
}

func NewLoggerNG() *LoggerNG {
	l := LoggerNG{
		curLevel:    INFO,
		exitOnError: false,
	}

	l.Logger.SetFlags(0)
	return &l
}

func (l *LoggerNG) Log(lvl LogLevel, format string, v ...interface{}) {
	if l.curLevel > lvl {
		return
	}

	if l.fn != nil {
		l.Logger.Printf("%s:%s", l.component+l.fn())
	}

	if lvl < INFO {
		l.Logger.Printf("-- %s --", logLevelString[lvl])
	}

	l.Logger.Printf(format, v...)

	switch lvl {
	case PANIC:
		panic("-- PANIC -- ")
	case FATAL:
		l.Println("-- FATAL --")
		os.Exit(1)
	default:
		l.Println("")
	}

}
