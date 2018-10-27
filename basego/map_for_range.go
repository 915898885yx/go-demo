package main

import (
	"fmt"
)

// 通过索引来操作切片，通过循环本身元素操作切片，真正的map元素无法得到初始化
func main() {
	items := make([]map[int]int, 5)
	fmt.Printf("%v\n", items)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
		items[i][3] = 2
	}
	fmt.Printf("version A:Value ofitems:%v\n", items)

	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 2
	}
	fmt.Printf("B is %v\n", items2)
}
