package utils

import (
	"errors"
	"io/fs"
	"os"
)

func ReadSlfFile(
	slfPath string,

	chunkSize int,

	onStat func(stats fs.FileInfo),

	onReadProgress func(perc float64),

	onReadComplete func(header SlfHeader),
) ([]SlfEntry, SlfHeader, fs.FileInfo, error) {
	if slfStats, err := os.Stat(slfPath); os.IsNotExist(err) {
		return []SlfEntry{}, SlfHeader{}, slfStats, err
	}

	var buffer, slfStats, err = ReadFileBuffer(slfPath, chunkSize, onStat, onReadProgress)

	if err != nil {
		return []SlfEntry{}, SlfHeader{}, slfStats, err
	}

	header, err := GetSlfHeader(buffer)

	if err != nil {
		return []SlfEntry{}, SlfHeader{}, slfStats, errors.New("Invalid slf file.")
	}

	onReadComplete(header)

	return GetSlfBufferEntries(buffer), header, slfStats, err
}
