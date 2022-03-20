package utils

import "regexp"

// Performs a global search for word characters on the given argument, removing invalid ones.
func SanitizeStringWord(value string) string {
	return regexp.MustCompile(`\W+`).ReplaceAllString(value, "")
}
