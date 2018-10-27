package main

import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 56, "bravo": 34, "charlie": 23, "delta": 87, "foxtrot": 12, "echo": 56, "juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("key:%v, Value:%v\n", k, v)
	}
	//按照key值排序
	//	keys := make([]string, len(barVal))
	//	i := 0
	//	for k, _ := range barVal {
	//		keys[i] = k
	//		i++
	//	}
	//	sort.Strings(keys)
	//	fmt.Println("sorted:\n")
	//	for _, k := range keys {
	//		fmt.Printf("key:%v, value:%v\n", k, barVal[k])
	//	}
	//按照value值排序
	value := make([]int, len(barVal))
	i := 0
	for _, v := range barVal {
		value[i] = v
		i++
	}
	sort.Ints(value)
	fmt.Printf("排序后的value值%v\n", value)
}
