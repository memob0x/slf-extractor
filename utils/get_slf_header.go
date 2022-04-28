package utils

import (
	"errors"
	"strings"
)

type SlfHeader struct {
	OriginalName string

	OriginalPath string
}

// Retrieves the main header informations from a buffer: the original slf file name and path.
func GetSlfHeader(buffer []byte) (SlfHeader, error) {
	var header SlfHeader
	var err error

	var pointer0 int = 0
	var pointer1 int = INT_BUFFER_STRING_LENGTH

	var originalSlfName string = SanitizeStringFilename(string(buffer[pointer0:pointer1]))

	var nameLowerCase string = strings.ToLower(originalSlfName)

	if !strings.HasSuffix(nameLowerCase, "slf") {
		err = errors.New("not a valid slf file")
	}

	header.OriginalName = originalSlfName

	pointer0 = pointer1
	pointer1 = pointer0 + INT_BUFFER_STRING_LENGTH

	header.OriginalPath = SanitizeStringFilename(string(buffer[pointer0:pointer1]))

	return header, err
}
