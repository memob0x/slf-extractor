package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeStringFilename(t *testing.T) {
	assert.Equal(t, "foobar.mp3", SanitizeStringFilename("foobar.mp3"), "should be the same")

	assert.Equal(t, "foobar.mp3", SanitizeStringFilename("::foobar$°.mp3$$$"), "should escape invalid characters")
}
