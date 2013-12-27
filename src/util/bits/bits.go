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
	var c int // c will be the number of zero bits on the right

	// Taken from bit twidling hacks stanford article
	b8 &= -signed(b8)
	if b8 {
		c--
	}
	if b8 & 0x0000FFFF {
		c -= 16
	}
	if b8 & 0x00FF00FF {
		c -= 8
	}
	if b8 & 0x0F0F0F0F {
		c -= 4
	}
	if b8 & 0x33333333 {
		c -= 2
	}
	if b8 & 0x55555555 {
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

func (b64 B64) IsSetBit(p int) bool {
	return (b64>>uint(p))&1 == 1
}

func (b64 B64) IsClearBit(p int) bool {
	return (b64>>uint(p))&1 == 0
}
