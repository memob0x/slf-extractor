package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatMs(t *testing.T) {
	assert.Equal(t, "1 milliseconds", FormatMs(1), "should return milliseconds formatted as milliseconds")

	assert.Equal(t, "1 seconds", FormatMs(1000), "should return milliseconds formatted as seconds")

	assert.Equal(t, "1 minutes", FormatMs(60000), "should return milliseconds formatted as minutes")

	assert.Equal(t, "1 hours", FormatMs(3600000), "should return milliseconds formatted as hours")

	assert.NotEqual(t, "1 days", FormatMs(86400000), "should not return milliseconds formatted as days (not supported)")
}
