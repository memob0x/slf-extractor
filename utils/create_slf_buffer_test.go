package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSlfBuffer(t *testing.T) {
	var buffer []byte = CreateSlfBuffer("TEST.SLF", "foobar/", []SlfEntry{
		{
			Name: "first.txt",

			Data: []byte("foo"),
		},
		{
			Name: "second.txt",

			Data: []byte("bar"),
		},
	})

	var header, _ = GetSlfHeader(buffer)

	assert.Equal(t, "TEST.SLF", header.OriginalName, "should be able to read original slf file name")
	assert.Equal(t, "foobar/", header.OriginalPath, "should be able to read original slf file path (relative do original installation path \"Data\" folder")

	var entries = GetSlfBufferEntries(buffer)

	var entriesCount int = len(entries)

	assert.Equal(t, 2, entriesCount, "TEST.slf should have 10 entries")

	for i := 0; i < entriesCount; i++ {
		var entry SlfEntry = entries[i]

		assert.IsTypef(t, string(""), entry.Name, "entry should have a string \"name\" member")

		assert.IsTypef(t, []byte{}, entry.Data, "entry should have a bytes array \"data\" member")

		if i == 0 {
			assert.Equal(t, "first.txt", entry.Name, "entry \"name\" should have the right content")

			assert.Equal(t, "foo", string(entry.Data), "entry \"data\" should have the right content")
		}

		if i == 1 {
			assert.Equal(t, "second.txt", entry.Name, "entry \"name\" should have the right content")

			assert.Equal(t, "bar", string(entry.Data), "entry \"data\" should have the right content")
		}
	}
}
