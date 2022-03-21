package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSlfBuffer(t *testing.T) {
	var buffer []byte = CreateSlfBuffer("TEST.SLF", "foobar/", []entryInformation{
		{
			name: "first.txt",

			data: []byte("foo"),
		},
		{
			name: "second.txt",

			data: []byte("bar"),
		},
	})

	var header, _ = GetSlfHeader(buffer)

	assert.Equal(t, "TEST.SLF", header[0], "should be able to read original slf file name")
	assert.Equal(t, "foobar/", header[1], "should be able to read original slf file path (relative do original installation path \"Data\" folder")

	var entries = GetSlfBufferEntries(buffer)

	var entriesCount int = len(entries)

	assert.Equal(t, 2, entriesCount, "TEST.slf should have 10 entries")

	for i := 0; i < entriesCount; i++ {
		var entry entryInformation = entries[i]

		assert.IsTypef(t, string(""), entry.name, "entry should have a string \"name\" member")

		assert.IsTypef(t, []byte{}, entry.data, "entry should have a bytes array \"data\" member")

		if i == 0 {
			assert.Equal(t, "first.txt", entry.name, "entry \"name\" should have the right content")

			assert.Equal(t, "foo", string(entry.data), "entry \"data\" should have the right content")
		}

		if i == 1 {
			assert.Equal(t, "second.txt", entry.name, "entry \"name\" should have the right content")

			assert.Equal(t, "bar", string(entry.data), "entry \"data\" should have the right content")
		}
	}
}
