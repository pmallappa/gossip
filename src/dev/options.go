package dev

import(
	"util"
)

func ParseFlags() (map[string]string, error) {
	m, e := util.ParseFlags(devflags)
	if e != nil {
		return m, e
	}

	for k, v := range m {
		switch k {
		case "char":
		default:
			v = v
		}
	}

	return m, e
}
