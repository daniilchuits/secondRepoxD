package main

import (
	"log"
	"os"
	"sync"
	"workers/output"
)

func main() {

	inputDir := os.Getenv("INPUT_DIR")
	outputPath := os.Getenv("OUTPUT_PATH")

	filesNames, err := output.ReadingInputDir(inputDir)
	if err != nil {
		log.Println("Error reading dir:", err)
		return
	}

	filesNames, err = output.CheckTxt(filesNames)
	if err != nil {
		log.Println(err)
		return
	}

	var wg sync.WaitGroup
	chWithData := make(chan output.Data)
	var wgWriter sync.WaitGroup
	wgWriter.Add(1)

	go func() {
		defer wgWriter.Done()
		if err = output.WriteToResult(outputPath, chWithData); err != nil {
			log.Println("Error to open result.txt:", err)
			return
		}
	}()

	for _, file := range filesNames {
		wg.Add(1)

		go func(fileAndPath output.NameAndPath) {
			defer wg.Done()

			data, err := output.ReadFile(file)
			if err != nil {
				log.Println("ReadFile err:", err)
				return
			}

			chWithData <- *data
		}(file)
	}

	wg.Wait()
	close(chWithData)

	wgWriter.Wait() // вообще на маленьких данных не обязательно делать wgWriter.Wait()
}
