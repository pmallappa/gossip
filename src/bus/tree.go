package bus

func (b *Bus) getDevice(addr uint64) (ReadWriteAll, uint64, error) {
	devlist := b.devices
	for len(devlist) > 0 {
		mid = len(devlist)/2
		d := devlist[mid]
		if addr >= d.start && addr <= d.end {
			return d.dr, addr - d.start, d.device
		}

		if addr > d.end {
			devlist = devlist[mid:len(devlist)]
		}
		if addr < d.start {
			devlist = devlist[0:mid-1]
		}
	}
	
	return nil, 0, errors.New("Illegal Address Access")
}

func (b *Bus) getDevice1(addr uint64) (ReadWriteAll, uint64, error) {
	// Since most read/writes are going to memory
	// We may need to check first for memory
	d := b
	for d != nil {
		if addr > d.device.end {
			d = d.Right
			continue
		}

 		if addr < b.device.start {
			d = d.Left
			continue
		}

		if addr >= d.device.start && addr < d.device.end {
			return b.device.dr, addr - d.device.start, nil
		}
	}
	return nil, 0, errors.New("Device Not found")
}

func (b *Bus) balance() {
}

func (b *Bus) sort() {

}

func (b *Bus) add(addr, size uint64, rw ReadWriteAll) error {
	for i, _ := range b.devices {
		if b.devices[i].start == addr {
			return errors.New("Device already exists")
	}
	b.devices = append(b.devices, &device{addr, size, addr+size, rw}
	return nil
}