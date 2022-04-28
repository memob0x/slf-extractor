package utils

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSlfHeader(t *testing.T) {
	WriteFile("bizz.slf", CreateSlfBuffer("original-name.slf", "./original/path", []SlfEntry{
		{
			Name: "first.txt",

			Data: []byte("foo"),
		},
		{
			Name: "second.txt",

			Data: []byte("bar"),
		},
	}))

	var buffer, _, _ = ReadFileBuffer("./bizz.slf", 8, func(stat fs.FileInfo) {}, func(_ float64) {})

	var header, _ = GetSlfHeader(buffer)

	assert.Equal(t, "original-name.slf", header.OriginalName, "should be able to read original slf file name")
	assert.Equal(t, "./original/path", header.OriginalPath, "should be able to read original slf file path (relative do original installation path \"Data\" folder")

	os.Remove("bizz.slf")
}
