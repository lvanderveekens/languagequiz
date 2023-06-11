package exercise

import (
	mystrings "languagequiz/utils/strings"
	"strings"
)

func normalizeAnswer(answer string) string {
	return strings.ToLower(strings.TrimSpace(mystrings.NormalizeApostrophes(answer)))
}
