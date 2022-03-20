package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSlfHeader(t *testing.T) {
	if !TEST_SLF_FILES {
		fmt.Printf("TEST_SLF_FILES env var is not set, skipping this test case \n")

		return
	}

	var buffer, _, _ = ReadFileBuffer("./test.slf", func(_ int) {})

	var header, _ = GetSlfHeader(buffer)

	assert.Equal(t, "AMBIENT.SLF", header[0], "should be able to read original slf file name")
	assert.Equal(t, "Ambient/", header[1], "should be able to read original slf file path (relative do original installation path \"Data\" folder")
}
