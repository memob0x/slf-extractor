package utils

import (
	"errors"
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

	onWriteComplete func(files []os.File),
) (fs.FileInfo, []os.File, SlfHeader, error) {
	if slfStats, err := os.Stat(slfPath); os.IsNotExist(err) {
		return slfStats, nil, SlfHeader{}, err
	}

	var buffer, slfStats, err = ReadFileBuffer(slfPath, chunkSize, onStat, onReadProgress)

	if err != nil {
		return slfStats, nil, SlfHeader{}, err
	}

	header, err := GetSlfHeader(buffer)

	if err != nil {
		return slfStats, nil, header, errors.New("Invalid slf file.")
	}

	onReadComplete(header)

	var entries []SlfEntry = GetSlfBufferEntries(buffer)

	var writtenFiles []os.File

	for i, j := 0, len(entries); i < j; i++ {
		var entry SlfEntry = entries[i]

		writtenFile, err := WriteFile(destinationPath+"/"+entry.name, entry.data)

		if err != nil {
			return slfStats, nil, header, err
		}

		onWriteProgress(writtenFile)

		writtenFiles = append(writtenFiles, *writtenFile)
	}

	onWriteComplete(writtenFiles)

	return slfStats, writtenFiles, header, err
}
