package logng

// Requirements,
//     Caller calls like
//                    <object>.Log(fmt, ...)
//                    <object>.Printf(fmt, ...)
//                    <object>.Log.V(INFO).Printf(fmt, ...)
//                    <object>.Info(string)
//                    <object>.Critical(string)
//                    <object>.Error(string)
//                    <object>.Panic(string)
// where <object> can be package name for logging, or components can embed the
// logger to log to their own file/telnet connection (helpful in debugging
// particular component)
//
// Additional information like componenet name, UUID etc can be added to logger
// along with Date/time info
// And allows objects to register a callback function to send more information
//
// "INFO", will print all the things the simulator is doing, used for debugging the simulator
// "DEBUG", allows to debug the Components
// "WARNING", is unexpected things, but simulator can still run
// "CRITICAL", is unexpected things, needing immediate attention, but simulator can still run
// "ERROR", unwanted thing happening, simulator can't run, waits for
//          debugger/cli to fix the condition, should print caller's location
// "FATAL", Unfavourable condition, simulator will exit, should print callers location
// By default LogLevel is set to "WARNING", All the levels above the current
// LogLevel will be logged, all logs below will be discarded
//
// -log level=INFO,out=file:<path>
// -log level=WARNING,out=tcp:localhost:2000
//          if 'out' is not given STDOUT is assumed
//
//
//
//

import (
	//"runtime"
	"flag"
	//"fmt"
	//"io"
	"log"
	"os"
	"strings"
	//"sync"
)

import (
//"util"
//"util/telnet"
)

type LogLevel uint8

const (
	INFO LogLevel = iota
	DEBUG
	WARNING
	CRITICAL
	ERROR
	FATAL
)

var levelToName = []string{
	INFO:     "INFO",
	DEBUG:    "DEBUG",
	WARNING:  "WARNING",
	CRITICAL: "CRITICAL",
	ERROR:    "ERROR",
	FATAL:    "FATAL",
}

var nameToLevel = map[string]LogLevel{
	"INFO":     INFO,
	"DEBUG":    DEBUG,
	"WARNING":  WARNING,
	"CRITICAL": CRITICAL,
	"ERROR":    ERROR,
	"FATAL":    FATAL,
}

// This one is for one without any object, but want to log,
// This will connect whatever is given on the commandline or to STDERR
var stderrargs string
var stderr LogNG

func (l *LogNG) SetLevel(lvl LogLevel) {
	l.level = lvl
}

func (l *LogNG) Set(str string) error {
	for _, val := range strings.Split(str, ",") {
		args := strings.Split(val, "=")
		switch args[0] {
		case "level":
			l.level = nameToLevel[strings.ToUpper(args[1])]
		case "out":
			// if strings.HasPrefix(args[1], "tcp:") ||
			// 	strings.HasPrefix(args[1], "udp:") {
			// 	// Open a telnet session,
			// 	return
			// }
			if strings.HasPrefix(args[1], "file:") {
				args[1] = args[1][5:]
				// Else its a file, open the file
				if writer, e := os.OpenFile(args[1], os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0640); e != nil {
					return e
				} else {
					l.log = log.New(writer, "", 0)
				}
			} else {
				// Store for whoever wants to use
				if l.args == "" {
					l.args += val
				} else {
					l.args += "," + val
				}
			}

		}
	}

	// If we are still not set
	l.log = log.New(os.Stderr, "", 0)

	return nil
}

func init() {
	flag.StringVar(&stderrargs, "log", "", "Set the log output and level"+
		"eg: -log level=<l>,out=file:<path>"+
		" OR -log level=<l>,out=<custom_arg>"+
		" valid <l>are INFO,DEBUG,WARNING,CRITICAL,FATAL")

	// If nothing is set, STDERR is opened

	// We dont want to be silent while starting, later the driver can set the
	// level to WARNING, to suppress the storm
	stderr.level = INFO
}

type LogNG struct {
	// level is the current level, checked everytime before logging
	level LogLevel

	// Use the actual logger, it has lot of feature, we dont want to re-invent
	// the wheel
	log *log.Logger

	// Extra args, that we were unable to Parse()
	args string
}

// Create new logger with given string, which is parsed using Set
func New(str string) *LogNG {
	l := &LogNG{
		level: WARNING,
	}
	l.Set(str)

	return l
}

func (l *LogNG) Args() string { return l.args }

func (l *LogNG) EnableDate() { l.log.SetFlags(log.Ldate) }

func (l *LogNG) Info(args ...interface{}) {
	if l.level > INFO {
		l.log.Print(args...)
	}
}

func (l *LogNG) Infoln(args ...interface{}) {
	if l.level <= INFO {
		l.log.Println(args...)
	}
}

func (l *LogNG) Infof(fmt string, args ...interface{}) {
	if l.level <= INFO {
		l.log.Printf(fmt, args...)
	}
}

func (l *LogNG) Warning(args ...interface{}) {
	if l.level <= WARNING {
		l.log.Print(args...)
	}
}

func (l *LogNG) Warningln(args ...interface{}) {
	if l.level <= WARNING {
		l.log.Println(args...)
	}
}

func (l *LogNG) Warningf(fmt string, args ...interface{}) {
	if l.level <= WARNING {
		l.log.Printf(fmt, args...)
	}
}

func (l *LogNG) Error(args ...interface{}) {
	if l.level <= ERROR {
		l.log.Fatal(args...)
	}
}

func (l *LogNG) Errorln(args ...interface{}) {
	if l.level <= ERROR {
		l.log.Fatalln(args...)
	}
}

func (l *LogNG) Errorf(fmt string, args ...interface{}) {
	if l.level <= ERROR {
		l.log.Fatalf(fmt, args...)
	}
}

func (l *LogNG) Fatal(args ...interface{}) {
	if l.level <= FATAL {
		f := l.log.Flags()
		l.log.SetFlags(log.Llongfile)
		l.log.Panic(args...)
		l.log.SetFlags(f)
	}
}

func (l *LogNG) Fatalln(args ...interface{}) {
	if l.level <= FATAL {
		// We want to report where the Fatal error happened
		f := l.log.Flags()
		l.log.SetFlags(log.Llongfile)
		l.log.Panicln(args...)
		l.log.SetFlags(f)
	}
}

func (l *LogNG) Fatalf(fmt string, args ...interface{}) {
	if l.level <= FATAL {
		f := l.log.Flags()
		l.log.SetFlags(log.Llongfile)
		l.log.Panicf(fmt, args...)
		l.log.SetFlags(f)
	}
}

func ParseLogger(s string) (*LogNG, error) {
	return nil, nil
}
