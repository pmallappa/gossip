package plat

import (
//"debug/elf"
)

import (
	"util/logng"
)

func zeroLoad() error {
	for addr, length := range zeroloads {
		curPlat.logger.LogLevel(logng.INFO,
			"Writing zeros to %x bytes @ %x", addr, length)
		buf := make([]byte, 1024)
		for length > 0 {
			if length > uint64(len(buf)) {
				if _, err := curPlat.busMain.WriteAt(buf, addr); err != nil {
					return err
				}
			} else {
				b := buf[:length]
				if _, err := curPlat.busMain.WriteAt(b, addr); err != nil {
					return err
				}
				// We have reached the end
				break
			}
			addr += uint64(len(buf))
			length -= uint64(len(buf))
		} // for length != 0
	} // for addr, length
	return nil
}
