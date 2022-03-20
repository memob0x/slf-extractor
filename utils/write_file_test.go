package utils

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteFile(t *testing.T) {
	var path string = "foo/bar/bizz.txt"

	var expectationFirstWrite string = "This is a string"

	WriteFile(path, []byte(expectationFirstWrite))

	resultFirstRead, _ := ioutil.ReadFile(path)

	assert.Equal(t, expectationFirstWrite, string(resultFirstRead), "should be able to write files in nested directories")

	var expectationOverwrite string = "This is a string 1111123"

	WriteFile(path, []byte(expectationOverwrite))

	resultOverwrittenRead, _ := ioutil.ReadFile(path)

	assert.Equal(t, expectationOverwrite, string(resultOverwrittenRead), "should be able to overwrite files in nested directories")

	os.RemoveAll("foo")
}
