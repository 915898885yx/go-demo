package main

import (
	"fmt"
	"math/rand"
	"time"
)

//rand.Float32()和rand.Float64返回介于[0.0,1.0)之间的伪随机数
//rand.Intn返回[0,n)之间的伪随机数
//使用seed(Value)函数提供伪随机数的生成种子，一般情况下会使用当前时间的纳秒级数字
func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d\n", a)
	}

	for i := 0; i < 5; i++ {
		r := rand.Intn(8)
		fmt.Printf("%d\n", r)
	}
	fmt.Println()

	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2.2f\n", 100*rand.Float32())
	}
}
