package utils

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractSlfEntriesTest(t *testing.T) {
	WriteFile("FOOBAR.SLF", CreateSlfBuffer("FOOBAR.SLF", ".", []SlfEntry{
		{
			Name: "first.txt",

			Data: []byte("foo"),
		},
		{
			Name: "second.txt",

			Data: []byte("bar"),
		},
	}))

	ExtractSlfEntries("./FOOBAR.SLF", "./", 8, func(stats fs.FileInfo) {}, func(perc float64) {}, func(header SlfHeader) {}, func(file *os.File) {}, func(files []*os.File) {})

	_, err0 := os.Stat("FOOBAR.SLF")

	assert.Nil(t, err0, "ok")

	_, err1 := os.Stat("first.txt")

	assert.Nil(t, err1, "ok")

	_, err2 := os.Stat("second.txt")

	assert.Nil(t, err2, "ok")

	os.Remove("FOOBAR.SLF")
	os.Remove("first.txt")
	os.Remove("second.txt")
}
