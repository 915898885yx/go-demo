package main

import (
	"fmt"
)

func getPrice() (int, string) {
	i := 2
	s := "goldfish"
	return i, s
}

func main() {
	quantity, price := getPrice()
	fmt.Printf("you won %v %v", quantity, price)
}
