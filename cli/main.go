package cli

import (
	"os"

	"github.com/memob0x/slf-exporter/utils"
)

func init() {
	var argsCount = len(os.Args)

	var slfPath string

	var destPath string = "output"

	if argsCount >= 2 {
		slfPath = os.Args[1]
	}

	if argsCount >= 3 {
		destPath = os.Args[2]
	}

	// TODO: printf some progress informations
	utils.ExtractSlfEntries(slfPath, destPath)
}
