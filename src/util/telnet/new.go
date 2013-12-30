package telnet

import (
//	"util"
)

func (t *telnetT) handleCMD() (c byte, again bool, err error) {

	if c, err = t.bufrd.ReadByte(); err != nil {
		return 0, false, err
	}

	//t.opts = opt
	if t.debug {
		//util.PrintMe()
		printcmd(c)
	}

	// In case Do/Dont/Will/Wont
	// Check we can satisfy the request
	// OR we initiate a negotiation
	switch c {
	case cmd_WILL:
		//t._do()
	case cmd_WONT:
		//t._dont()
	case cmd_DO:
		//t._will()
	case cmd_DONT:
		//t._wont()

	case cmd_IAC:
		again = true
		err = nil
		return
	default:
		t.cmd = c
		again = true
		err = t.__execCMD(c)
		return
	}

	return

}
