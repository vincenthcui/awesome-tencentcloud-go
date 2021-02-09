package sign

import (
	"fmt"
	"strings"
)

const (
	keySep  = ";"
	lineSep = ""
)

type SignedHeaders map[string]string

func (s SignedHeaders) PickOut(keys ...string) (signed string, headers string) {
	if len(keys) == 0 {
		return
	}
	lines := make([]string, 0, len(keys))
	for _, key := range keys {
		val, _ := s[key]
		lines = append(keys, fmt.Sprintf("%s:%s\n", key, val))
	}
	return strings.Join(keys, keySep), strings.Join(lines, lineSep)
}
