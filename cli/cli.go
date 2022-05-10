package cli

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/memob0x/slf-extractor/utils"
)

func ParseArgs() (string, string) {
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

func Launch() {
	var slfPath, destPath = ParseArgs()

	var stat, _, _, err = utils.ExtractSlfEntries(
		slfPath,

		destPath,

		1048576, // 1MB

		func(stats fs.FileInfo) {
			fmt.Printf("Start extracting %v.\n", stats.Name())
		},

		func(perc float64) {
			fmt.Printf("\rReading progress: %v%%\r", utils.FormatFloat(perc, 2))
		},

		func(header utils.SlfHeader) {
			fmt.Printf("Read complete %v.\n", header)
		},

		func(file *os.File) {
			var s, _ = file.Stat()

			fmt.Printf("Writing %v (%v)\n", file.Name(), utils.FormatBytes(s.Size()))
		},

		func(files []*os.File) {
			fmt.Printf("\nDone.\n")
		},
	)

	if stat == nil {
		log.Fatalf("The given file doesn't exist or it's invalid.\n")
	}

	if err != nil {
		log.Fatalf(err.Error())
	}
}
