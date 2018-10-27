package main

import (
	"fmt"
)

func main() {
	// 生命一个长度为2切片的语法
	var cheeses = make([]string, 2)
	cheeses[0] = "slice's first element"
	cheeses[1] = "slice's second element"
	fmt.Println(cheeses)
	// adding element to slice
	cheeses = append(cheeses, "slice's three element")
	cheeses = append(cheeses, "1", "2", "3")
	fmt.Println(cheeses)
}
