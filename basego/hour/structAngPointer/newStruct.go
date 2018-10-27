package main

import (
	"fmt"
)

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	m := new(Movie)
	m.Name = "mrwang"
	m.Rating = 32.8
	fmt.Printf("%+v\n", m)
}
