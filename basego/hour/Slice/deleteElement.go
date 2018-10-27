package main

import (
	"fmt"
)

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "first element"
	cheeses[1] = "second element"
	cheeses[2] = "three element"
	fmt.Println(len(cheeses))
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println(len(cheeses))
	fmt.Println(cheeses)
}
