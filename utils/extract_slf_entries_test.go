package utils

import (
	"fmt"
	"testing"
)

func TestExtractSlfEntriesTest(t *testing.T) {
	if !TEST_SLF_FILES {
		fmt.Printf("TEST_SLF_FILES env var is not set, skipping this test case \n")

		return
	}

	ExtractSlfEntries("./test.slf", "./output")

	// TODO: add test cases
}
