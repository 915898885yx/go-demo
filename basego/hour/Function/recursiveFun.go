package main

import (
	"fmt"
)

func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("i'm full! i've eaten %d\n", eaten)
		return eaten
	}
	fmt.Printf("i'm still hugry! i've eat %d\n", eaten)
	return feedMe(portion, eaten)
}

func main() {
	fmt.Println(feedMe(1, 0))
}
