package util

import (
	"bufio"
	"log"
	"os"
)

type FileReader struct {
	Filename string
	Contents []string
}

func New(path string) FileReader {
	newReader := FileReader{path, nil}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newReader.Contents = append(newReader.Contents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return newReader
}
