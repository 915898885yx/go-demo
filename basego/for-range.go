package main

import (
	"fmt"
)

func main() {
	seasons := []string{"spring", "summer", "autumn", "winter"}
	for ix, season := range seasons {
		fmt.Print(ix, season, "\n")
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
	// 如果只要索引，可以忽略第二个变量
	for ix := range seasons {
		fmt.Printf("%d", ix)
	}
	// 0123

}
