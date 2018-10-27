package main

import (
	"fmt"
)

var (
	barVal = map[string]int{"alpha": 56, "bravo": 34, "charlie": 23, "delta": 87, "foxtrot": 12, "echo": 56, "juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	invMap := make(map[int]string, len(barVal))

	for k, v := range barVal {
		invMap[v] = k
	}
	fmt.Println(invMap)
}
