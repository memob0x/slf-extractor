package utils

import (
	"io/fs"
	"log"
	"os"
)

// TODO: better errors management
func ExtractSlfEntries(slfPath string, destinationPath string, chunkSize int, onProgress func(perc float64)) (fs.FileInfo, error) {
	if _, err := os.Stat(slfPath); os.IsNotExist(err) {
		log.Fatal("Slf file does not exist.")
	}

	var buffer, stats, err = ReadFileBuffer(slfPath, chunkSize, onProgress)

	var entries []entryInformation = GetSlfBufferEntries(buffer)

	for i, j := 0, len(entries); i < j; i++ {
		var entry entryInformation = entries[i]

		_, err := WriteFile(destinationPath+"/"+entry.name, entry.data)

		if err != nil {
			return stats, err
		}
	}

	return stats, err
}
