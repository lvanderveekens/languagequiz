package strings

import (
	"strings"
)

func NormalizeApostrophes(s string) string {
	var normalized strings.Builder
	for _, c := range s {
		if c == '’' {
			normalized.WriteRune('\'')
		} else {
			normalized.WriteRune(c)
		}
	}
	return normalized.String()
}
