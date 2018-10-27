package main

import (
	"fmt"
)

func main() {
	var chssese = make(map[string]int)
	chssese["cook"] = 32
	chssese["num"] = 55
	fmt.Println(chssese)
	delete(chssese, "cook")
	fmt.Println(chssese)
}
