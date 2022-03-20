package utils

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLittleEndianUnsignedInt32(t *testing.T) {
	var buffer []byte = make([]byte, 4)

	binary.LittleEndian.PutUint32(buffer, 100)

	assert.Equal(t, 100, GetLittleEndianUnsignedInt32(buffer, 0, len(buffer)), "should be able to extract little endian unsigned int 32 values from a bytes array")
}
