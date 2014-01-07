package bits

type (
	B8  uint8
	B16 uint16
	B32 uint32
	B64 uint64
)

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
	return int(b8 & 1 << uint(b)), nil
}

func (b8 B8) Set(b int) error {
	if b > 7 {
		return E_InvalidPosition
	}
	b8 |= 1 << uint(b)
	return nil
}

func (b8 B8) Clear(b int) error {
	if b > 7 {
		return E_InvalidPosition
	}
	b8 &= ^(1 << uint(b))
	return nil
}

func (b8 B8) IsSet(b int) (bool, error) {
	if b > 7 {
		return false, E_InvalidPosition
	}
	return b8&1<<uint(b) == 1, nil
}

func (b8 B8) IsClear(b int) bool {
	if b > 7 {
		return false
	}
	return b8&1<<uint(b) == 0
}

// Count Leading zeros, from MSB
func (b8 B8) Clz() int {
	// c will be the number of zero bits on the right
	var c int = 8

	// Taken from bit twidling hacks stanford article
	b8 &= B8(-b8)
	if b8 != 0 {
		c--
	}
	if b8&0x0F != 0 {
		c -= 4
	}
	if b8&0x33 != 0 {
		c -= 2
	}
	if b8&0x55 != 0 {
		c -= 1
	}

	return 0
}

// Count Trailing zeros, from LSB
func (b8 B8) Ctz() int {
	return 0
}

/*
 * #####    ##   ####
 * ##  ##   ##  ##
 * #####    ##  ####',
 * ##  ##   ##  ##  ##
 * #####    ##  ######
 */

func (b16 B16) Set(b int) {
	if b > 16 {
		return
	}
	b16 |= 1 << uint(b)
}

func (b16 B16) Clear(b int) {
	if b > 16 {
		return
	}
	b16 &= ^(1 << uint(b))

}

func (b16 B16) IsSet(b int) bool {
	if b > 16 {
		return false
	}
	return b16&1<<uint(b) == 1
}

func (b16 B16) IsClear(b int) bool {
	if b > 16 {
		return false
	}
	return b16&1<<uint(b) == 0
}

// Count Leading zeros, from MSB
func (b B16) Clz() int {
	var c int = 16 // c will be the number of zero bits on the right

	// Taken from bit twidling hacks stanford article
	b &= B16(-b)
	if b != 0 {
		c--
	}
	if b&0x00FF != 0 {
		c -= 8
	}
	if b&0x0F0F != 0 {
		c -= 4
	}
	if b&0x3333 != 0 {
		c -= 2
	}
	if b&0x5555 != 0 {
		c -= 1
	}

	return 0
}

/*
 * #####    #####    ###
 * ##  ##      ##  ##  ##
 * #####   ###',     ##
 * ##  ##     ##   ##
 * #####  #####  ######
 */
func (b32 B32) Bit(p int) int {
	if p > 31 {
		panic("Invalid bit position")
	}
	return int((b32 >> uint(p)) & 1)
}

func (b32 B32) IsSet(p int) bool {
	if p > 31 {
		panic("")
	}
	return (b32>>uint(p))&1 == 1
}

func (b32 B32) IsClear(p int) bool {
	if p > 31 {
		panic("")
	}
	return (b32>>uint(p))&1 == 0
}

func (b32 B32) Bits(high, low uint8) uint32 {
	return uint32((b32 >> low) & (^(1 << (high - low + 1)) - 1))
}

// Count Leading zeros, from MSB
func (b B32) Clz() int {
	var c int = 32 // c will be the number of zero bits on the right

	// Taken from bit twidling hacks stanford article
	b &= B32(-b)
	if b != 0 {
		c--
	}
	if b&0x0000FFFF != 0 {
		c -= 16
	}
	if b&0x00FF00FF != 0 {
		c -= 8
	}
	if b&0x0F0F0F0F != 0 {
		c -= 4
	}
	if b&0x33333333 != 0 {
		c -= 2
	}
	if b&0x55555555 != 0 {
		c -= 1
	}

	return 0
}

/*
 *     #####  ####   ##
 *    ##  ## ##     ##   ##
 *   #####  ####', ##   ##
 *  ##  ## ##  ## #########
 * #####  ######      ##
 */

func (b64 B64) Bit(p int) int {
	if p > 63 {
		return -1
	}
	return int((b64 >> uint(p)) & 1)
}

func (b64 B64) Bits(high, low int) uint64 {
	return uint64((b64 >> uint(low)) & (^(1 << uint(high-low+1)) - 1))
}

func (b64 B64) IsSet(p int) bool {
	return (b64>>uint(p))&1 == 1
}

func (b64 B64) IsClear(p int) bool {
	return (b64>>uint(p))&1 == 0
}

// Count Leading zeros, from MSB
func (b B64) Clz() int {
	var c int = 16 // c will be the number of zero bits on the right

	// Taken from bit twidling hacks stanford article
	b &= B64(-b)
	if b != 0 {
		c--
	}
	if b&0x00000000FFFFFFFF != 0 {
		c -= 32
	}
	if b&0x0000FFFF0000FFFF != 0 {
		c -= 16
	}
	if b&0x00FF00FF00FF00FF != 0 {
		c -= 8
	}
	if b&0x0F0F0F0F0F0F0F0F != 0 {
		c -= 4
	}
	if b&0x3333333333333333 != 0 {
		c -= 2
	}
	if b&0x5555555555555555 != 0 {
		c -= 1
	}

	return 0
}
