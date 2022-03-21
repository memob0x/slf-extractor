package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSlfHeader(t *testing.T) {
	WriteFile("bizz.slf", CreateSlfBuffer("original-name.slf", "./original/path", []entryInformation{
		{
			name: "first.txt",

			data: []byte("foo"),
		},
		{
			name: "second.txt",

			data: []byte("bar"),
		},
	}))

	var buffer, _, _ = ReadFileBuffer("./bizz.slf", 8, func(_ float64) {})

	var header, _ = GetSlfHeader(buffer)

	assert.Equal(t, "original-name.slf", header[0], "should be able to read original slf file name")
	assert.Equal(t, "./original/path", header[1], "should be able to read original slf file path (relative do original installation path \"Data\" folder")

	os.Remove("bizz.slf")
}
