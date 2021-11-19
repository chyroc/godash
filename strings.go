package godash

import (
	"strings"
	"unicode/utf8"
)

func CountPre(s, substr string) int {
	if len(substr) == 0 {
		return utf8.RuneCountInString(s) + 1
	}
	n := 0
	for {
		i := strings.Index(s, substr)
		if i == -1 {
			break
		}
		if i != len(substr)-1 {
			break
		}
		n++
		s = s[i+1:]
	}
	return n
}
