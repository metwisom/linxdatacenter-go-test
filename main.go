package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	PathError   = "Data path not specified. ex.: ./data/1.json\n"
	FileNF      = "The specified file does not exist\n"
	FormatError = "Unknown data format %s\n"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal(PathError)
	}

	var filePath = os.Args[1]

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		log.Fatal(FileNF)
	}

	var ext = filepath.Ext(filePath)

	var dearest, popular *Line

	if ext == ".json" {
		dearest, popular = readJsonFile(filePath)
	} else if ext == ".csv" {
		dearest, popular = readCsvFile(filePath)
	} else {
		log.Fatal(fmt.Sprintf(FormatError, ext))
	}

	fmt.Println("Dearest: " + dearest.Name)
	fmt.Println("Popular: " + popular.Name)

}
