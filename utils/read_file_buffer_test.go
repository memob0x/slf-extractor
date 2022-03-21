package utils

import (
	"bytes"
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileBuffer(t *testing.T) {
	var percentageFormer float64 = -1

	WriteFile("test.txt", []byte("foobar"))

	var buffer, _, _ = ReadFileBuffer(
		"./test.txt",

		8,

		func(stat fs.FileInfo) {},

		func(percentage float64) {
			assert.IsTypef(t, float64(0), percentage, "should return an int")

			assert.GreaterOrEqual(t, percentage, percentageFormer, "percentage should increment")

			percentageFormer = percentage
		},
	)

	assert.Equal(t, "foobar", bytes.NewBuffer(buffer).String(), "should be able to read files returning them as buffer")

	os.Remove("test.txt")
}
