// compress包读取压缩文件
package main

import(
	"fmt"
	"bufio"
	"os"
	"compress/gzip"
)

func main() {
	fName := "MyFile.gz"
	
	var r *bufio.Reader
	fi, err := os.Open(fName)
	
	if err != nil {
		fmt.Fprint(os.Stderr, "%v, Cant't open %s: error:%s\n", os.Args[0].fName)
		os.Exit(1)
	}
	fz. err:= gzip.NewReader(fi)
	
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}
	
	for {
		line, err := r.ReadString('\n')
		if err := nil {
			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}