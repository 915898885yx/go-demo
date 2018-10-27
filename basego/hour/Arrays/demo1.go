package main

import (
	"fmt"
)

func main() {
	var cheeses [2]string
	cheeses[0] = "array's first element"
	cheeses[1] = "array's second element"
	// 长度一旦声明，不能随意更改长度
	//	cheeses[2] = "last element"
	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	fmt.Println(cheeses)
}
