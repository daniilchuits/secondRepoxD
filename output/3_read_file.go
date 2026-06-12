package output

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Data struct {
	FilePath string
	FileName string
	Lines    int
	Words    int
	Chars    int
}

func ReadFile(nameAndPath NameAndPath) (*Data, error) {

	bytes, err := os.ReadFile(nameAndPath.Path)
	if err != nil {
		log.Printf("Error reading file %s: %s\n", nameAndPath.Name, err)
		return nil, fmt.Errorf("Error reading file %s: %s\n", nameAndPath.Name, err)
	}

	txt := string(bytes)
	lines := strings.Split(txt, "\n")
	words := strings.Fields(txt)

	return &Data{
		FilePath: nameAndPath.Path,
		FileName: nameAndPath.Name,
		Lines:    len(lines),
		Words:    len(words),
		Chars:    len(txt),
	}, nil

}
