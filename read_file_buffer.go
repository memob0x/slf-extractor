package main

import (
	"bytes"
	"os"
)

const macChunkSize int = 4

func ReadFileBuffer(path string, onProgress func(progress int)) (*bytes.Buffer, error) {
	var output = bytes.NewBuffer(make([]byte, 0))
	var err error

	stats, err := os.Stat(path)

	var fileSize int = int(stats.Size())

	if err != nil {
		return output, err
	}

	file, err := os.Open(path)

	if err != nil {
		return output, err
	}

	defer file.Close()

	var part []byte = make([]byte, macChunkSize)

	var readLengthTotal int = 0

	for {
		readLengthPart, err := file.Read(part)

		readLengthTotal += readLengthPart

		var percentage int = readLengthTotal / fileSize * 100

		onProgress(percentage)

		if err != nil {
			break
		}

		output.Write(part[:readLengthPart])
	}

	return output, err
}
