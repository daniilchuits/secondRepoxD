package output

import (
	"fmt"
	"strings"
)

func CheckTxt(namesAndPaths []NameAndPath) ([]NameAndPath, error) {

	var checkedNames []NameAndPath

	for _, nameAndPath := range namesAndPaths {

		if ok := strings.HasSuffix(nameAndPath.Name, ".txt"); ok {
			checkedNames = append(checkedNames, nameAndPath)
		}
	}

	if len(checkedNames) == 0 {
		return nil, fmt.Errorf("No '.txt' files in dir")
	}

	return checkedNames, nil
}
