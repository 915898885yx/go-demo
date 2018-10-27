package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Index返回字符串str在字符串s中的索引（str第一个字符的索引），
//-1表示字符串s不包含字符串str
//strings.Index(s, str string) int

//LastIndex返回字符串str在字符串s中最后出现位置的索引（str第一个字符串的所用），
//-1表示不包含

//如果ch是非ASCII编码的字符，建议用以下对字符串定位
//strings.IndexRune(s string, ch int) int

//Repeat用于重复count次字符串s并返回一个新的字符串
//strings.Repeat(s,count int) string

//ToLower将字符串中的Unicode字符串全部转换为小写
//strings.ToLower(s) string

//ToUpper将字符串中的Unicode字符全部转换为大写
//strings.ToUpper(s) string

//strings.TrimSpace(s)剔除开头和结尾的空格
//strings.Trim(s, "cut")开头和结尾的cut去掉
//去掉开头或者结尾TrimLeft、TrimRight

//strings.Fields(s)将会利用1个或多个空白符号分割为若干小块并返回一个slice

//Join用于将元素类型为string的slice使用分隔符拼接为一个字符串
//strings.Join(sl []string, sep string)
func main() {
	var origS string = "Hi there"
	var newS string
	newS = strings.Repeat(origS, 3)
	fmt.Printf("the new repeated string is:%s\n", newS)

	var lower string
	lower = strings.ToUpper(origS)
	fmt.Printf("大写为：%s\n", lower)
	lower = strings.ToLower(origS)
	fmt.Printf("小写为：%s\n", lower)

	var str string = "cuta123456bcut"
	var cutstr = strings.Trim(str, "cut")
	fmt.Printf("去掉cut的值为：%s\n", cutstr)
	sliceStr := strings.Split(cutstr, "，")
	fmt.Printf("切割后的slice为:%v\n", sliceStr)
	fmt.Printf("slice的值：%d", len(sliceStr[0]))

	//	str := "The quick brown fox jumps over"
	//	str1 := strings.Split(str, "|")
	//	fmt.Printf("split:%v\n", str1)
	//	str3 := strings.Join(str1, ";")
	//	fmt.Printf("join后的值为%s", str3)

	// Read()从[]type 中读取内容
	// ReadByte()和ReadRune()从字符串中读取下一个byte和rune

	//字符相关类型的转换都是通过.strconv包实现
	//任何类型转为字符串总是成功的

	//strconv.Itoa(i int) string返回数字i表示的字符串类型的十进制数字
	//strconv.FormatFloat(f float64, fmt byte, prec int, biteSize int)string
	// 其中 fmt 表示格式（其值可以是 'b' 、 'e' 、 'f' 或 'g' ）， prec 表示精度， bitSize  则使用 32 表示 float32，用 64 表示 float64。

}
