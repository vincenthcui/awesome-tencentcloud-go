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

func (s SignedHeaders) PickOut(fields ...string) (signedFields string, signedHeaders string) {
	if len(fields) == 0 {
		return
	}

	lines := make([]string, 0, len(fields))
	for _, field := range fields {
		val, _ := s[field]
		field = strings.ToLower(field)
		lines = append(fields, fmt.Sprintf("%s:%s\n", field, val))
	}

	// 参数要求小写
	for idx, key := range fields {
		fields[idx] = strings.ToLower(key)
	}

	return strings.Join(fields, keySep), strings.Join(lines, lineSep)
}
