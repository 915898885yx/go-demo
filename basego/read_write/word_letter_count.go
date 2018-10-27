package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines int

func main() {
	nrchars, nrwords, nrlines = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input out S:")
	for {
		input, err := inputReader.ReadString('\n')
		//		fmt.Println('1111111')
		if err != nil {
			fmt.Printf("An error occurred:%s\n", err)
		}
		if input == "S\r\n" {
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words:%d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}
		Counters(input)
	}
}

func Counters(input string) {
	nrchars += len(input) - 2
	nrwords += len(strings.Fields(input))
	nrlines++
}
