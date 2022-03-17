package main

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
)

func readCsvFile(filePath string) (*Line, *Line) {
	f := OpenReader(filePath)
	r := csv.NewReader(f)
	r.Comma = ';'
	if len(os.Args) == 3 {
		r.Comma = ([]rune(os.Args[2]))[0]
	}
	r.Read()

	data, err := readData(r)
	if err != nil {
		if err.Error() == "ParseError" {
			log.Fatal("CSV format is wront please use second arg for set comma ex.: ./app $path$ \",\"")
		}
		log.Fatal("Some error on csv reader")
	}
	var dearest Line = *data
	var popular Line = *data

	for {
		Rec, err := readData(r)
		if err != nil {
			break
		}
		if dearest.Price < Rec.Price {
			dearest = *Rec
		}
		if popular.Rating < Rec.Rating {
			popular = *Rec
		}
	}

	return &dearest, &popular
}

func readData(r *csv.Reader) (*Line, error) {
	record, err := r.Read()
	if err == io.EOF {
		return nil, err
	}
	if err != nil {
		log.Fatal(err)
	}
	if len(record) != 3 {
		return nil, errors.New("ParseError")
	}
	var Rec Line
	Rec.Parse(record)
	return &Rec, nil
}
