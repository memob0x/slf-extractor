package cli

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/memob0x/slf-exporter/utils"
)

func onStat(stats fs.FileInfo) {
	fmt.Printf("Start extracting.\n")
}

func onReadComplete(header utils.SlfHeader) {
	fmt.Printf("%v.\n", header)
}

func onReadProgress(percentage float64) {
	fmt.Printf("\rExtracting: %v%%\r", utils.FormatFloat(percentage, 2))
}

func onWriteComplete(files []os.File) {
	fmt.Printf("\nDone.\n")
}

func onWriteProgress(file *os.File) {
	fmt.Printf("Writing.%v\n", file.Name())
}

func parseArgs() (string, string) {
	var argsCount = len(os.Args)

	var slfPath string

	var destPath string = "output"

	if argsCount >= 2 {
		slfPath = os.Args[1]
	} else {
		log.Fatalf("No slf file specified, aborting.\n")
	}

	if argsCount >= 3 {
		destPath = os.Args[2]
	} else {
		fmt.Printf("No output folder specified, fallback to \"%v\".\n", destPath)
	}

	return slfPath, destPath
}

func init() {
	var slfPath, destPath = parseArgs()

	var stat, _, _, err = utils.ExtractSlfEntries(
		slfPath,

		destPath,

		1048576, // 1MB

		onStat,

		onReadProgress,

		onReadComplete,

		onWriteProgress,

		onWriteComplete,
	)

	if stat == nil {
		log.Fatalf("The given file doesn't exist or it's invalid.\n")
	}

	if err != nil {
		log.Fatalf(err.Error())
	}
}
