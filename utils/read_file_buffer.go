package utils

import (
	"io"
	"io/fs"
	"os"
)

func ReadFileBuffer(path string, chunkSize int, onStat func(stat fs.FileInfo), onProgress func(percentage float64)) ([]byte, os.FileInfo, error) {
	var buffer = make([]byte, 0)
	var err error

	stats, err := os.Stat(path)

	var fileSize int = int(stats.Size())

	if err != nil {
		return buffer, stats, err
	}

	onStat(stats)

	file, err := os.Open(path)

	if err != nil {
		return buffer, nil, err
	}

	defer file.Close()

	var chunk []byte = make([]byte, chunkSize)

	var readLengthTotal int = 0

	for {
		readLengthPart, readErr := file.Read(chunk)

		if readErr != io.EOF {
			err = readErr
		}

		if readErr != nil {
			break
		}

		buffer = append(buffer, chunk[:readLengthPart]...)

		readLengthTotal += readLengthPart

		onProgress(float64(readLengthTotal) / float64(fileSize) * 100)
	}

	return buffer, stats, err
}
