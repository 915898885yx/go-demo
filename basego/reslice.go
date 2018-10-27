package main

import (
	"fmt"
)

func main() {
	//reslice

	slice1 := make([]int, 0, 10)
	fmt.Printf("初始化的切片为%v\n", slice1)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("the length of slice is %d\n", len(slice1))
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("重组的切片为%v\n", slice1)

	// append_slice
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_from)
	fmt.Printf("cpoied %v\n", n)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}
