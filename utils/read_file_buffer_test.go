package utils

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileBuffer(t *testing.T) {
	var percentageFormer int = -1

	var buffer, _, _ = ReadFileBuffer("./test.txt", func(percentage int) {
		assert.IsTypef(t, int(0), percentage, "should return an int")

		assert.GreaterOrEqual(t, percentage, percentageFormer, "percentage should increment")

		percentageFormer = percentage
	})

	assert.Equal(t, "foobar", bytes.NewBuffer(buffer).String(), "should be able to read files returning them as buffer")
}