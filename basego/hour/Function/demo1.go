package main

import (
	"fmt"
)

func isEven(i int) bool {
	return i%2 == 0
}

func main() {
	var i int = 5
	fmt.Println(isEven(i))
}
