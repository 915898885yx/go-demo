package main

import (
	"fmt"
)

const greeting string = "Hello world"

func main() {
	greeting = "goodbye, cruel world"
	fmt.Println(greeting)
}
