// Various Bit operations.
// static declarations are used to speed up the process,

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

func (b8 B8) Bits(high, low int) uint8 {
	h := uint(high)
	l := uint(low)
	if h < l {
		h, l = l, h
	}
	return uint8((b8 >> l) & (^(1 << (h - l + 1)) - 1))
}

func (b8 B8) Bit(b int) int {
	return int(b8 & 1 << uint(b))
}

func (b8 B8) Set(b int) {
	b8 |= 1 << uint(b)
}

func (b8 B8) Clear(b int) {
	b8 &= ^(1 << uint(b))
}

func (b8 B8) IsSet(b int) bool {
	return b8&1<<uint(b) == 1
}

func (b8 B8) IsClear(b int) bool {
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

func Bit8(v uint8, p int) int {
	return B8(v).Bit(p)
}

func Bits8(v uint8, high, low int) uint8 {
	return B8(v).Bits(high, low)
}

func IsSet8(v uint8, p int) bool {
	return B8(v).IsSet(p)
}

func IsClear8(v uint8, p int) bool {
	return B8(v).IsSet(p)
}

func Clz8(v uint8) int {
	return B8(v).Clz()
}

/*
 *    #####    ##   ####
 *   ##  ##   ##  ##
 *  #####    ##  ####',
 * ##  ##   ##  ##  ##
 *#####    ##  ######
 */
func (b16 B16) Bits(high, low int) uint16 {
	h := uint(high)
	l := uint(low)
	if h < l {
		h, l = l, h
	}

	return uint16((b16 >> l) & (^(1 << (h - l + 1)) - 1))
}

func (b16 B16) Bit(p int) int {
	return int(b16 & 1 << uint(p))
}

func (b16 B16) Set(b int) {
	b16 |= 1 << uint(b)
}

func (b16 B16) Clear(b int) {
	b16 &= ^(1 << uint(b))
}

func (b16 B16) IsSet(b int) bool {
	return b16&1<<uint(b) == 1
}

func (b16 B16) IsClear(b int) bool {
	return b16&1<<uint(b) == 0
}

func Bit16(v uint16, p int) int {
	return B16(v).Bit(p)
}

func Bits16(v uint16, high, low int) uint16 {
	return B16(v).Bits(high, low)
}

func IsSet16(v uint16, p int) bool {
	return B16(v).IsSet(p)
}

func IsClear16(v uint16, p int) bool {
	return B16(v).IsSet(p)
}

func Clz16(v uint16) int {
	return B16(v).Clz()
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
 *    #####  #####   ###
 *   ##  ##     ## ##  ##
 *   #####   ###',    ##
 *  ##  ##   ##     ##
 * #####  #####   ######
 */
func (b32 B32) Bit(p int) int {
	return int((b32 >> uint(p)) & 1)
}

func (b32 B32) IsSet(p int) bool {
	return (b32>>uint(p))&1 == 1
}

func (b32 B32) IsClear(p int) bool {
	return (b32>>uint(p))&1 == 0
}

func (b32 B32) Bits(high, low int) uint32 {
	h := uint(high)
	l := uint(low)
	if h < l {
		h, l = l, h
	}
	return uint32((b32 >> l) & (^(1 << (h - l + 1)) - 1))
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

func Bit32(v uint32, p int) int {
	return B32(v).Bit(p)
}

func Bits32(v uint32, high, low int) uint32 {
	return B32(v).Bits(high, low)
}

func IsSet32(v uint32, p int) bool {
	return B32(v).IsSet(p)
}

func IsClear32(v uint32, p int) bool {
	return B32(v).IsSet(p)
}

func Clz32(v uint32) int {
	return B32(v).Clz()
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
	h := uint(high)
	l := uint(low)
	if h < l {
		h, l = l, h
	}
	return uint64((b64 >> l) & (^(1 << (h - l + 1)) - 1))
}

func (b64 B64) IsSet(p int) bool {
	return (b64>>uint(p))&1 == 1
}

func (b64 B64) IsClear(p int) bool {
	return (b64>>uint(p))&1 == 0
}

// Count Leading zeros, from MSB
func (b B64) Clz() int {
	var c int = 64 // c will be the number of zero bits on the right

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

func Bit64(v uint64, p int) int {
	return B64(v).Bit(p)
}

func Bits64(v uint64, high, low int) uint64 {
	return B64(v).Bits(high, low)
}

func IsSet64(v uint64, p int) bool {
	return B64(v).IsSet(p)
}

func IsClear64(v uint64, p int) bool {
	return B64(v).IsSet(p)
}

func Clz64(v uint64) int {
	return B64(v).Clz()
}
