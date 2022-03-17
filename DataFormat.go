package main

import "strconv"

type Line struct {
	Name   string `json:"product"`
	Price  int    `json:"price"`
	Rating int    `json:"rating"`
}

func (t *Line) Parse(in []string) {
	t.Name = in[0]
	t.Price, _ = strconv.Atoi(in[1])
	t.Rating, _ = strconv.Atoi(in[2])
}
