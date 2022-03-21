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
	} else {
		fmt.Printf("No slf file specified, aborting.\n")

		return
	}

	if argsCount >= 3 {
		destPath = os.Args[2]
	} else {
		fmt.Printf("No output folder specified, fallback to \"%v\".\n", destPath)
	}

	// TODO: output more informations about the extraction progress

	utils.ExtractSlfEntries(slfPath, destPath, func(percentage int) {
		fmt.Printf("%d%%\r", percentage)
	})

	fmt.Printf("Done.\n")
}
