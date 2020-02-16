package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello,世界")

	//var s ,sep string

	var s = "aaa"
	var sep = "bbb"
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], "  "))
	fmt.Println(os.Args[0])
}

//注意 { 必须和 func 在同一行 不能单独成行
//go run

//go build

//格式化
//在对应目录下执行 go fmt

//os包 提供一些函数和变量
