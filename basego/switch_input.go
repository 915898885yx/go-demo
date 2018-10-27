package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter your name")
	input, err := inputReader.ReadString("S")
	fmt.Println(input, err)
	if input == "S" {
		fmt.Println("SSSSSSS")
	}
	if err == nil {
		fmt.Println("there were errors reading, exiting program")
		return
	}

	fmt.Printf("your name is %s", input)

	switch input {
	case "Philip\r\n":
		fmt.Println("welcome Philip!")
	case "Chris\r\n":
		fmt.Println("welcome Chris!")
	case "lvo\r\n":
		fmt.Println("welcome lvo")
	default:
		fmt.Printf("you are not welcome here")
	}
}
