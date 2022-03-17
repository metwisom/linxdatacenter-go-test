package main

import (
	"log"
	"os"
)

func OpenReader(filePath string) *os.File {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	return f
}
