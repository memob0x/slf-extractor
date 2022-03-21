package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractSlfEntriesTest(t *testing.T) {
	WriteFile("FOOBAR.SLF", CreateSlfBuffer("FOOBAR.SLF", ".", []entryInformation{
		{
			name: "first.txt",

			data: []byte("foo"),
		},
		{
			name: "second.txt",

			data: []byte("bar"),
		},
	}))

	ExtractSlfEntries("./FOOBAR.SLF", "./", func(perc int) {})

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
