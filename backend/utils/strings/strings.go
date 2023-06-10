package strings

import "regexp"

var punctRegex = regexp.MustCompile(`[[:punct:]]`)

func RemovePunctuation(s string) string {
	return punctRegex.ReplaceAllString(s, "")
}
