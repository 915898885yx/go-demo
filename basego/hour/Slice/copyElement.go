package main

import (
	"fmt"
)

func main() {
	var cheeses = make([]string, 3)
	cheeses[0] = "first Element"
	cheeses[1] = "second Element"
	cheeses[2] = "three element"
	var smellCheese = make([]string, 2)
	// 将切片cheeses中的元素复制到smellCheeses，但是长度
	// 不会超过smellCheese的长度
	//	copy(smellCheese, cheeses)
	copy(smellCheese, cheeses[1:])
	fmt.Println(smellCheese)
}
