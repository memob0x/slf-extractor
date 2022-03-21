package utils

import (
	"log"
	"os"
)

// TODO: better errors management
func ExtractSlfEntries(slfPath string, destinationPath string, onProgress func(percentage int)) {
	if _, err := os.Stat(slfPath); os.IsNotExist(err) {
		log.Fatal("Slf file does not exist.")
	}

	var buffer, _, _ = ReadFileBuffer(slfPath, onProgress)

	var entries []entryInformation = GetSlfBufferEntries(buffer)

	for i, j := 0, len(entries); i < j; i++ {
		var entry entryInformation = entries[i]

		WriteFile(destinationPath+"/"+entry.name, entry.data)
	}
}
