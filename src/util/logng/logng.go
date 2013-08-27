package log

import (
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
	logInfo LogLevel = iota
	logWarning
	logError
	logPanic
)

type LoggerNG struct {
	log.Logger
	str         string
	logType     LoggerType
	level       LogLevel
	curlevel    LogLevel
	exitOnError bool
}

func (log *LoggerNG) SetLevel(l LogLevel) {
	log.level = l
}

func ParseLogger(s string) (*LoggerNG, error) {
	loggerstr := strings.SplitN(s, ":", 2)
	var logger *LoggerNG = &LoggerNG{
		logType: TcpUdp,
	}
	switch strings.ToLower(loggerstr[0]) {
	case "", "tcp", "udp":
		logger.logType = TcpUdp
	case "file":
		fallthrough
	default:
		logger.logType = File
	}
	logger.str = loggerstr[1]
	return logger, nil
}

func (l *LoggerNG) InitLogger(ls string) (*log.Logger, error) {
	var logwriter io.Writer
	var e error
	switch l.logType {
	case TcpUdp:
		if logwriter, e = telnet.Start(ls); e != nil {
			return nil, e
		}
	case File:
		if logwriter, e = os.OpenFile(ls, os.O_WRONLY|os.O_CREATE,
			0640); e != nil {
			return nil, e
		}
	}

	return log.New(logwriter, "", 0), e
}

func NewLoggerNG() *LoggerNG {
	return &LoggerNG{
		level:       logInfo,
		exitOnError: false,
	}
}

func (l *LoggerNG) Printf(format string, v ...interface{}) {
	l.Logger.Printf(format, v...)
}
