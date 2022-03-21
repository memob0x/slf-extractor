package utils

import (
	"os"
	"path/filepath"
)

func WriteFile(name string, data []byte) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(name), 0770); err != nil {
		return nil, err
	}

	file, err := os.Create(name)

	defer file.Close()

	file.Write(data)

	return file, err
}
