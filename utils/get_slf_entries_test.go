package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSlfEntries(t *testing.T) {
	WriteFile("bazz.slf", CreateSlfBuffer("bazz.slf", ".", []entryInformation{
		{
			name: "first.txt",

			data: []byte("foo"),
		},
		{
			name: "second.txt",

			data: []byte("bar"),
		},
	}))

	var buffer, _, _ = ReadFileBuffer("./bazz.slf", func(_ int) {})

	var entries = GetSlfBufferEntries(buffer)

	var entriesCount int = len(entries)

	assert.Equal(t, 2, entriesCount, "bazz.slf should have 2 entries")

	for i := 0; i < entriesCount; i++ {
		var entry entryInformation = entries[i]

		assert.IsTypef(t, string(""), entry.name, "should have a string \"name\" member")

		assert.IsTypef(t, []byte{}, entry.data, "should have a bytes array \"data\" member")
	}

	os.Remove("bazz.slf")
}
