package plat

import (
	"debug/elf"
	"errors"
	"fmt"
)

import (
//	"bus"
)

type PlatELFLoader interface {
	GetELFMachine() []elf.Machine
	//GetABI() elf.OSABI
	GetELFClass() []elf.Class
}

func isClassSupported(c elf.Class) bool {
	for _, class := range curPlat.GetELFClass() {
		if c == class {
			return true
		}
	}
	return false
}

func isMachineSupported(m elf.Machine) bool {
	for _, machine := range curPlat.GetELFMachine() {
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
		} else {
			return errors.New("Elf File problem")
		}
		defer file.Close() // Close after loaded
	}
	return nil
}
