package utils

import (
	"os"
)

func GetBufferChunkSize(fileSize int) int {
	return 700
}

func ReadFileBuffer(path string, onProgress func(progress int)) ([]byte, os.FileInfo, error) {
	var buffer = make([]byte, 0)
	var err error

	stats, err := os.Stat(path)

	var fileSize int = int(stats.Size())

	if err != nil {
		return buffer, stats, err
	}

	file, err := os.Open(path)

	if err != nil {
		return buffer, nil, err
	}

	defer file.Close()

	var chunk []byte = make([]byte, GetBufferChunkSize(fileSize))

	var readLengthTotal int = 0

	for {
		readLengthPart, err := file.Read(chunk)

		readLengthTotal += readLengthPart

		if err != nil {
			break
		}

		var percentage int = readLengthTotal / fileSize * 100

		onProgress(percentage)

		buffer = append(buffer, chunk[:readLengthPart]...)
	}

	return buffer, stats, err
}
