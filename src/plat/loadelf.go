package plat

import (
	//	"bytes"
	"debug/elf"
	"errors"
	"fmt"
)

import (
//"bus"
)

type PlatELFLoader interface {
	GetELFMachine() []elf.Machine
	//GetABI() elf.OSABI
	GetELFClass() []elf.Class
}

func isClassSupported(c elf.Class) bool {
	for _, class := range curPlatform.GetELFClass() {
		if c == class {
			return true
		}
	}
	return false
}

func isMachineSupported(m elf.Machine) bool {
	for _, machine := range curPlatform.GetELFMachine() {
		if m == machine {
			return true
		}
	}
	return false
}

// for every file being loaded
//    check if MACHINE is supported, we cant load an ARM binary to MIPS
//    check if CLASS is supported
//    check if OSABI is supported

func loadELF(files []string) error {
	for _, filename := range files {
		var file *elf.File
		if file, err := elf.Open(filename); err != nil {
			if !isMachineSupported(file.Machine) {
				return fmt.Errorf("Machine not supported: %v", file.Machine)
			}
			if !isClassSupported(file.Class) {
				return fmt.Errorf("Class not supported: %v", file.Class)
			}
			for _, proghdr := range file.Progs {
				// load each program header
				if proghdr.Flags&(elf.PF_R|elf.PF_X|elf.PF_W) != 0 {
					// MemSz will be aligned so we no need to write '0'
					// at end of FileSz bytes
					p := make([]byte, proghdr.Memsz)
					if _, e := proghdr.ReadAt(p, int64(proghdr.Off)); e != nil {
						if _, err := curPlat.busMain.WriteAt(p, proghdr.Vaddr); err != nil {
							return err
						}
					}
				}
			}
		} else {
			return errors.New("Elf File problem")
		}
		defer file.Close() // Close after loaded
	}
	return nil
}
