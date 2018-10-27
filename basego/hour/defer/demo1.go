package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("run a function1 complete")
	defer fmt.Println("run a function2 complete")
	defer fmt.Println("run a function3 complete")
	fmt.Println("hello world")
	return
}
