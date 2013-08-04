package util

import (
	"fmt"
	"io"
	"log"
	"runtime"
)

var dbg int

type debug struct {
	subsys string
	uid    uint16
	log    *log.Logger
	level  int
}

func PrintMe() {
	pc, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s=>%20s:%d\n", runtime.FuncForPC(pc).Name(), file, line)
}

func printMyFunc(s string, n int) {
	pc, _, line, _ := runtime.Caller(n)
	fmt.Printf("%s=>%s:%d\n", s, runtime.FuncForPC(pc).Name(), line)
}

func Entered() {
	if dbg > 0 {
		printMyFunc("Entering", 2)
	}
}

func Exiting() {
	if dbg > 0 {
		printMyFunc("Exiting", 2)
	}
}

func DebugInit(subsys string, uid int, w io.Writer) {

}

func SetDebugLevel(n int) {
	dbg = n
}

func GetDebugLeve() int {
	return dbg
}
