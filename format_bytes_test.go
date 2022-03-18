package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesMs(t *testing.T) {
	assert.Equal(t, "999 kB", FormatBytes(999), "should be able to format kB")

	assert.Equal(t, "987.7 MB", FormatBytes(987654321), "should be able to format MB")

	assert.Equal(t, "9.2 EB", FormatBytes(9223372036854775807), "should be able to format max")
}
