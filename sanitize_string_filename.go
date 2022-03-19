package main

import (
	"fmt"
	"strings"
)

// Gets the given file name with all invalid characters removed.
func SanitizeStringFilename(value string) string {
	var parts []string = strings.Split(value, ".")

	var extension string = parts[len(parts)-1]

	// FIXME: allow "/"
	return fmt.Sprintf("%v.%v", SanitizeStringWord(strings.Replace(value, extension, "", -1)), SanitizeStringWord(extension))
}
