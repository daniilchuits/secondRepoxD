package output

import (
	"fmt"
	"log"
	"os"
)

func WriteToResult(outputPath string, ch chan Data) error {

	f, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for data := range ch {

		wr := fmt.Sprintf(
			"File: %s\nLines: %d\nWords: %d\nChars: %d\n\n",
			data.FileName, data.Lines, data.Words, data.Chars,
		)

		_, err = f.WriteString(wr)
		if err != nil {
			log.Printf("Error writing data about %s: %v",
				data.FileName, err,
			)
		}
	}
	return nil
}
