package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4}
	for i, n := range numbers {
		fmt.Println("i is", i)
		fmt.Println("n is", n)
	}
}
