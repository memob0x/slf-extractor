package utils

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSlfEntries(t *testing.T) {
	if os.Getenv("TEST_SLF_FILES") != "1" {
		fmt.Printf("TEST_SLF_FILES env var is not set, skipping this test case \n")

		return
	}

	var buffer, _, _ = ReadFileBuffer("./test.slf", func(_ int) {})

	var entries = GetSlfBufferEntries(buffer)

	var entriesCount int = len(entries)

	assert.Equal(t, 10, entriesCount, "test.slf should have 10 entries")

	for i := 0; i < entriesCount; i++ {
		var entry entryInformation = entries[i]

		assert.IsTypef(t, string(""), entry.name, "should have a string \"name\" member")

		assert.IsTypef(t, []byte{}, entry.data, "should have a bytes array \"data\" member")
	}
}
