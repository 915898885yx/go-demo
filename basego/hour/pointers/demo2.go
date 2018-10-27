package main

import (
	"fmt"
)

func showMeoryAddress(x *int) {
	fmt.Println(x)
	return
}

func main() {
	i := 1
	fmt.Println(&i)
	showMeoryAddress(&i)
}
