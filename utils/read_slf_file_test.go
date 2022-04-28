package utils

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadSlfFile(t *testing.T) {
	WriteFile("bazz.slf", CreateSlfBuffer("bazz.slf", ".", []SlfEntry{
		{
			Name: "first.txt",

			Data: []byte("foo"),
		},
		{
			Name: "second.txt",

			Data: []byte("bar"),
		},
	}))

	var entries, _, _, _ = ReadSlfFile("./bazz.slf", 8, func(stats fs.FileInfo) {}, func(perc float64) {}, func(header SlfHeader) {})

	var entriesCount int = len(entries)

	assert.Equal(t, 2, entriesCount, "bazz.slf should have 2 entries")

	for i := 0; i < entriesCount; i++ {
		var entry SlfEntry = entries[i]

		assert.IsTypef(t, string(""), entry.Name, "should have a string \"name\" member")

		assert.IsTypef(t, []byte{}, entry.Data, "should have a bytes array \"data\" member")
	}

	os.Remove("bazz.slf")
}
