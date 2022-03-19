package main

import (
	"regexp"
)

// Gets the given file name with all invalid characters removed.
func SanitizeStringFilename(value string) string {
	var sanitized string = regexp.MustCompile(`[^A-z0-9_\-\/\.]`).ReplaceAllString(value, "")

	return regexp.MustCompile(`\\`).ReplaceAllString(sanitized, "/")
}
