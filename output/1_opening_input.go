package output

import (
	"log"
	"os"
	"path/filepath"
)

type NameAndPath struct {
	Name string
	Path string
}

func ReadingInputDir(inputDir string) ([]NameAndPath, error) {

	files, err := os.ReadDir(inputDir)
	if err != nil {
		log.Println("Err reading dir:", err)
		return nil, err
	}

	var namesAndPaths []NameAndPath

	for _, file := range files {
		namesAndPaths = append(namesAndPaths, NameAndPath{
			Name: file.Name(),
			Path: filepath.Join(inputDir, file.Name()),
		})
	}

	return namesAndPaths, nil
}
