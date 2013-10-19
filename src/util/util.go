package util

var BtoI = map[bool]int{
	false: 0,
	true:  1,
}

var BtoU8 = map[bool]uint8{
	false: 0,
	true:  1,
}

func Btou(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
