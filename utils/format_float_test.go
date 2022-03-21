package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FormatFloatTest(t *testing.T) {
	assert.Equal(t, 1.1, FormatFloat(1.111111111111111111, 1), "should be able to format float leaving 1 decimal")
	assert.Equal(t, 1.11, FormatFloat(1.111111111111111111, 2), "should be able to format float leaving 2 decimals")
}
