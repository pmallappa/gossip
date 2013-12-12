package bits

type Bit32 uint32

func (b Bit32) Bit(p uint8) int8 {
	if p > 31 {
		return -1
	}
	return int8((b >> p) & 1)
}

func (b Bit32) IsSetBit(p uint8) bool {
	return (b>>p)&1 == 1
}

func (b Bit32) IsClearBit(p uint8) bool {
	return (b>>p)&1 == 0
}

func (b Bit32) Bits(high, low uint8) int32 {
	if high > 31 || high < low {
		return -1
	}

	return int32((b >> low) & (^(1 << (high - low + 1)) - 1))
}

type Bit64 uint64

func (b Bit64) Bit(p uint8) int8 {
	if p > 31 {
		return -1
	}
	return (b >> p) & 1
}

func (b Bit64) Bits(high, low uint8) int64 {
	if high > 31 || high < low {
		return -1
	}

	return (b >> low) & (^(1 << (high - low + 1)) - 1)
}
