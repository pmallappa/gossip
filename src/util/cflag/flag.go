package cflag

import (
	//	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Parser interface {
	Parse() error
}

// KVPairs is comma separated key=value pairs
type StrFlag struct {
	Str     string
	KVPairs map[string]*KVFlag
}

func NewStrFlag(s string) *StrFlag {
	return &StrFlag{
		Str:     s,
		KVPairs: make(map[string]*KVFlag),
	}
}

func (sf *StrFlag) Set(str string) (e error) {
	sf.Str = str
	sf.Parse()
	return
}

func (sf *StrFlag) String() string {
	return sf.Str
}

type KVFlag struct {
	Name      string
	Default   interface{}
	ShortName string
	Desc      string // Description
	Value     interface{}
}

func NewKVFlag(name, desc string, def interface{}) *KVFlag {
	return &KVFlag{
		Name:    name,
		Desc:    desc,
		Default: def,
		Value:   def,
	}
}

func (kv *KVFlag) String() string {
	return fmt.Sprintf("Name %s, Default: %v, Description:%s, Value:%v",
		kv.Name,
		kv.Default,
		kv.Desc, kv.Value)
}

func (kv *KVFlag) Parse() error {
	return nil
}

func parseOne(kvp map[string]*KVFlag, s string) (err error) {

	split := strings.SplitN(s, "=", 2)

	if len(split) != 2 {
		return
	}
	k := kvp[split[0]]

	if k == nil {
		return fmt.Errorf("Not found")
	}

	// Confirmation
	if k.Name != split[0] {
		return fmt.Errorf("Unbelievable")
	}

	switch k.Default.(type) {
	case int8, int16, int32, int64, int:
		k.Value, err = strconv.ParseInt(split[1], 0, 64)

	case uint8, uint16, uint32, uint64, uint:
		k.Value, err = strconv.ParseUint(split[1], 0, 64)

	case float32, float64:
		k.Value, err = strconv.ParseFloat(split[1], 64)

	case bool:
		switch val := strings.ToLower(split[1]); val {
		case "yes", "true", "on", "1", "t":
			k.Value = true
		case "no", "false", "off", "0", "f":
			k.Value = false
		default:
			err = fmt.Errorf("Unfavourable value %s\n", split[1])
		}

	case string:

		k.Value = split[1]
	default:
		err = k.Parse()

	}

	fmt.Printf("Parsed %q\n", k)
	return
}

// Real parse function, for each of key=value pairs,
func (s *StrFlag) Parse() (k *KVFlag, e error) {
	str := s.Str
	split := strings.Split(str, ",")

	for i, kv := range split {
		if e = parseOne(s.KVPairs, split[i]); e != nil {
			return s.KVPairs[kv], e
		}
		fmt.Printf("%d\n", i)
	}
	return
}

func (s *StrFlag) Add(kvf *KVFlag) {
	s.KVPairs[kvf.Name] = kvf
	fmt.Printf("Added %q\n", kvf)
}
