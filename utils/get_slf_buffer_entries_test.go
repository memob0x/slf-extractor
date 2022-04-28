package utils

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSlfBufferEntries(t *testing.T) {
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

	var buffer, _, _ = ReadFileBuffer("./bazz.slf", 8, func(stat fs.FileInfo) {}, func(_ float64) {})

	var entries = GetSlfBufferEntries(buffer)

	var entriesCount int = len(entries)

	assert.Equal(t, 2, entriesCount, "bazz.slf should have 2 entries")

	for i := 0; i < entriesCount; i++ {
		var entry SlfEntry = entries[i]

		assert.IsTypef(t, string(""), entry.Name, "should have a string \"name\" member")

		assert.IsTypef(t, []byte{}, entry.Data, "should have a bytes array \"data\" member")
	}

	os.Remove("bazz.slf")
}
