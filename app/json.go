package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func readJsonFile(filePath string) (*Line, *Line) {
	f := OpenReader(filePath)
	file, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("Unable to parse file as JSON for "+filePath, err)
	}
	var records []Line
	json.Unmarshal(file, &records)
	var dearest = records[0]
	var popular = records[0]
	for i := 1; i < len(records); i++ {
		if records[i].Price > dearest.Price {
			dearest = records[i]
		}
		if records[i].Rating > popular.Rating {
			popular = records[i]
		}
	}
	return &dearest, &popular
}
