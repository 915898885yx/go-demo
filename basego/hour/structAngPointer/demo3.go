package main

import (
	"fmt"
)

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	m := Movie{
		Name:   "mrwang",
		Rating: 32,
	}
	fmt.Printf("%+v\n", m)
}
