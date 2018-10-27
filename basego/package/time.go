package main

import (
	"fmt"
	"time"
)

//Duration 类型表示两个连续时刻所相差的纳秒数
var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
}
