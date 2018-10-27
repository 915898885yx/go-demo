package main

import (
	"fmt"
	"strings"
)

//HasPrefix判断字符串s是否以prefix开头
//strings.HasPrefix(s, prefix string) bool

//HasSuffix判断字符串s是否以suffix结尾
//strings.HasSuffix(s, suffix string) bool

//Contains判断字符串s是否包含substr
//strings.Contains(s, substr sting) bool
func main() {
	var str = "this is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}
