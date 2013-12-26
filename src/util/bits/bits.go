package bits

type B8 uint8
type B16 uint16
type B32 uint32
type B64 uint64

type biterror uint8

const (
	E_InvalidPosition biterror = iota
)

var errormsg = []string{
	E_InvalidPosition: "Invalid Position",
}

func (b biterror) Error() string {
	return errormsg[b]
}

/*
 * #####    ####
 * ##  ##  ##  ##
 * #####   ',##',
 * ##  ##  ##  ##
 * #####    ####
 */

func (b8 B8) Bit(b int) (int, error) {
	if b > 7 {
		return 0, E_InvalidPosition
	}
	return b8 & 1 << b
}
func (b8 B8) Set(b int) error {
	if b > 7 {
		return E_InvalidPosition
	}
	b8 |= 1 << b
	return nil
}

func (b8 B8) Clear(b int) error {
	if b > 7 {
		return E_InvalidPosition
	}
	b8 &= ^(1 << b)
	return nil
}

func (b8 B8) IsSet(b int) (bool, error) {
	if b > 7 {
		return false, E_InvalidPosition
	}
	return b8&1<<b == 1, nil
}

func (b8 B8) IsClear(b int) bool {
	if b > 7 {
		return false, E_InvalidPosition
	}
	return b8&1<<b == 0
}

// Count Leading zeros, from MSB
func (b8 B8) Clz() int {
}

// Count Trailing zeros, from LSB
func (b8 B8) Ctz() int {

}

/*
 * #####    ##   ####
 * ##  ##   ##  ##
 * #####    ##  ####',
 * ##  ##   ##  ##  ##
 * #####    ##  ######
 */

func (b16 B16) Set(b int) error {
	if b > 16 {
		return E_InvalidPosition
	}
	b16 |= 1 << b
}

func (b16 B16) Clear(b int) error {
	if b > 16 {
		return E_InvalidPosition
	}
	b16 &= ^(1 << b)

}

func (b16 B16) IsSet(b int) bool {
	if b > 16 {
		return E_InvalidPosition
	}
	return b16&1<<b == 1
}

func (b16 B16) IsClear(b int) bool {
	if b > 16 {
		return E_InvalidPosition
	}
	return b16&1<<b == 0
}

/*
 * #####   ######    ###
 * ##  ##      ##  ##  ##
 * #####    ###',     ##
 * ##  ##      ##   ##
 * #####   ###### #######
 */
func (b32 Bit32) Bit(p int) int {
	if p > 31 {
		panic()
	}
	return (b32 >> p) & 1
}

func (b32 Bit32) IsSet(p int) bool {
	if p > 31 {
		panic()
	}
	return (b32>>p)&1 == 1
}

func (b32 Bit32) IsClear(p int) bool {
	if p > 31 {
		panic()
	}

	return (b32>>p)&1 == 0
}

func (b32 Bit32) Bits(high, low uint8) uint32 {
	if p > 31 {
		panic
	}

	return (b32 >> low) & (^(1 << (high - low + 1)) - 1)
}

/*
 * #####      ####  ##
 * ##  ##    ##     ##  ##
 * #####    ####',  ##  ##
 * ##  ##  ##  ##   ######
 * #####   ######       ##
 */

func (b64 Bit64) Bit(p int) int {
	if p > 63 {
		return -1
	}
	return (b >> p) & 1
}

func (b64 Bit64) Bits(high, low int) (uint64, error) {
	if high < low || high > 64 {
		return 0, E_InvalidPosition
	}
	return (b >> low) & (^(1 << (high - low + 1)) - 1), nil
}

func (b64 Bit64) IsSetBit(p int) bool {
	if p > 64 {
		return 0, E_InvalidPosition
	}
	return (b64>>p)&1 == 1
}

func (b64 Bit64) IsClearBit(p int) bool {
	if p > 64 {
		return 0, E_InvalidPosition
	}

	return (b64>>p)&1 == 0
}
