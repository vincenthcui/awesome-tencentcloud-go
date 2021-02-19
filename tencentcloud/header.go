package tencentcloud

import (
	"fmt"
	"strings"
)

const (
	keySep  = ";"
	lineSep = ""
)

func signHeaders(headers map[string]string, fields ...string) (signedFields string, signedHeaders string) {
	if len(fields) == 0 {
		return
	}

	lines := make([]string, 0, len(fields))
	for _, field := range fields {
		val, _ := headers[field]
		field = strings.ToLower(field)
		lines = append(lines, fmt.Sprintf("%s:%s\n", field, val))
	}

	// 参数要求小写
	for idx, key := range fields {
		fields[idx] = strings.ToLower(key)
	}

	return strings.Join(fields, keySep), strings.Join(lines, lineSep)
}

func toHttpHeader(headers map[string]string) map[string][]string {
	header := make(map[string][]string, 0)
	for k, v := range headers {
		header[k] = []string{v}
	}
	return header
}
