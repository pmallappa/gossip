package bits

// 1. Test Set bit
// 2. Test Clear bits
// 3. Test random bit mask
// 4. Count leading zeros
// 5. Count leading ones

type b8testpair struct {
	i      B8
	expset int
	expctz int
}

var testsB8 = []testpair{
	{0, 0, 8},
	{0x1, 1, 0},
	{0x2, 1, 1},
	{0x3, 2, 0},
	{0x4, 1, 2},
	{0x5, 2, 0},
	{0x6, 2, 1},
	{0x7, 3, 0},
}

func Test_B8(t *testing.T) {

}

type b16testpair struct {
	i      B16
	expset int
	expctz int
}

var testsB16 = []testpair{
	{0x00ff, 8, 1},
	{0x0100, 1, 8},
	{0x0500, 2, 8},
	{0xff00, 8, 0},
}

func Test_B16(t *testing.T) {

}

type b16testpair struct {
	i      B32
	expset int
	expctz int
}

var testsB32 = []testpair{
	{0x10, 1, 4},
	{0x50, 2, 4},
	{0xff, 8, 0},
}

func Test_B32(t *testing.T) {

}

type b16testpair struct {
	i      B64
	expset int
	expctz int
}

var testsB64 = []testpair{
	{0x10, 1, 4},
	{0x50, 2, 4},
	{0xff, 8, 0},
}

func Test_B64(t *testing.T) {

}
