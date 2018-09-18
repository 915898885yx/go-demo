package main

import(
	"fmt"
	"strings"
)

func main() {
	var string = "this is an example of a string"
	fmt.Printf("T/F does the string \"%s"have prefix %s?",str,"Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}