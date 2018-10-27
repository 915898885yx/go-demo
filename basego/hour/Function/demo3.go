package main

import (
	"fmt"
)

func sunNumbers(numbers ...int) int {
	total := 0
	for ap, number := range numbers {
		fmt.Println(ap)
		total += number
	}
	return total
}

func main() {
	result := sunNumbers(1, 1, 3, 4)
	fmt.Printf("the result is %v", result)
}
