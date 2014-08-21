package bus

import (
	"errors"
)

func (b *bus) getDevice(addr uint64) (ReadWriterAll, uint64, error) {
	devlist := b.devices
	for len(devlist) > 0 {
		mid := len(devlist) / 2
		d := devlist[mid]
		if addr >= d.start && addr <= d.end {
			return d.dr, addr - d.start, nil
		}

		if addr > d.end {
			devlist = devlist[mid:len(devlist)]
		}
		if addr < d.start {
			devlist = devlist[0 : mid-1]
		}
	}

	return nil, 0, errors.New("Illegal Address Access")
}

func (b *bus) balance() {
}

func (b *bus) sort() {

}

func (b *bus) add(addr, size uint64, rw ReadWriterAll) error {
	for i, _ := range b.devices {
		if b.devices[i].start == addr {
			return errors.New("Device already exists")
		}
	}
	b.devices = append(b.devices, &device{addr, size, addr + size, rw})
	b.sort()
	return nil
}
