package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeStringWord(t *testing.T) {
	assert.Equal(t, "foobar", SanitizeStringWord("foobar"), "should be the same")

	assert.Equal(t, "foobar", SanitizeStringWord("::foobar$°"), "should escape invalid characters")
}
