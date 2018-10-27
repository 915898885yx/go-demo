package main

import (
	"fmt"
)

func main() {
	// create a new map
	var players = make(map[string]int)
	players["cook"] = 32
	players["stock"] = 45
	fmt.Println(players)
}
