package tencentcloud

import "strings"

const (
	eol = "\n" // end of line
)

func joinLines(lines ...string) string {
	return strings.Join(lines, eol)
}
