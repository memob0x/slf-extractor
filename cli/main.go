package cli

import (
	"fmt"
	"os"

	"github.com/memob0x/slf-exporter/utils"
)

func init() {
	var argsCount = len(os.Args)

	var slfPath string

	var destPath string = "output"

	if argsCount >= 2 {
		slfPath = os.Args[1]

		fmt.Printf("Extraction of \"%v\" started.\n", slfPath)
	} else {
		fmt.Printf("No slf file specified, aborting.\n")

		return
	}

	if argsCount >= 3 {
		destPath = os.Args[2]
	} else {
		fmt.Printf("No output folder specified, fallback to \"%v\".\n", destPath)
	}

	// TODO: printf slf file size
	// TODO: printf slf read time
	// TODO: printf original slf name and path

	utils.ExtractSlfEntries(
		slfPath,

		destPath,

		1048576, // 1MB

		func(percentage float64) {
			fmt.Printf("Extraction: %v%%\r", utils.FormatFloat(percentage, 2))
		},
	)

	// TODO: printf extracted files total size
	// TODO: printf extracted files write time

	fmt.Printf("Done.\n")
}
