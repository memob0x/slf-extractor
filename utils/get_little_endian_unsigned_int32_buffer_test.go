package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLittleEndianUnsignedInt32Buffer(t *testing.T) {
	assert.IsTypef(t, []byte{}, GetLittleEndianUnsignedInt32Buffer(123, 4), "Should be able to convert an int to little endian unsigned int buffer (valid type)")

	assert.Equal(t, 123, GetLittleEndianUnsignedInt32Int(GetLittleEndianUnsignedInt32Buffer(123, 4), 0, 4), "Should be able to convert an int to little endian unsigned int buffer (valid value)")
}
