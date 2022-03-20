package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteFile(name string, data []byte) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(name), 0770); err != nil {
		return nil, err
	}

	file, err := os.Create(name)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	file.Write(data)

	return file, err
}
