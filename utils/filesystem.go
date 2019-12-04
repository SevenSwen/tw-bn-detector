package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetOutDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	dirPath := dir + "/out"

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.Mkdir(dirPath, os.ModeDir|os.ModePerm)
	}

	return dirPath
}
