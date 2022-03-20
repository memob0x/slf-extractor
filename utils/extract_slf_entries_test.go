package utils

import (
	"fmt"
	"os"
	"testing"
)

func TestExtractSlfEntriesTest(t *testing.T) {
	if os.Getenv("TEST_SLF_FILES") != "1" {
		fmt.Printf("TEST_SLF_FILES env var is not set, skipping this test case \n")

		return
	}

	ExtractSlfEntries("./test.slf", "./output")

	// TODO: add test cases
}
