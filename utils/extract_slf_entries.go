package utils

import (
	"io/fs"
	"os"
)

func ExtractSlfEntries(
	slfPath string,

	destinationPath string,

	chunkSize int,

	onStat func(stats fs.FileInfo),

	onReadProgress func(perc float64),

	onReadComplete func(header SlfHeader),

	onWriteProgress func(file *os.File),

	onWriteComplete func(files []*os.File),
) (fs.FileInfo, []*os.File, SlfHeader, error) {
	var entries, header, slfStats, err = ReadSlfFile(slfPath, chunkSize, onStat, onReadProgress, onReadComplete)

	var writtenFiles []*os.File

	for i, j := 0, len(entries); i < j; i++ {
		var entry SlfEntry = entries[i]

		writtenFile, err := WriteFile(destinationPath+"/"+entry.Name, entry.Data)

		if err != nil || writtenFile == nil {
			return slfStats, nil, header, err
		}

		onWriteProgress(writtenFile)

		writtenFiles = append(writtenFiles, writtenFile)
	}

	onWriteComplete(writtenFiles)

	return slfStats, writtenFiles, header, err
}
