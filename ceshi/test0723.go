package main

import "fmt"

func main() {
	hello := "hello"
	fmt.Println([]rune(hello))
	fmt.Println([]byte(hello))

	//每个英文字母占一个字节

	nihao := "你好"
	fmt.Println([]rune(nihao))
	fmt.Println([]byte(nihao))

	//每个中文字占三个字节

	//截取
	//左闭又开
	fmt.Println(hello[:2])
	//fmt.Println(nihao[:2]) //出现特殊符号 ？

	//将中文 [] rune 转换成 unicode 码点， 再利用 string 转化回去
	newNiHao := "你好大家"
	newNiHao2 := []rune(newNiHao)
	fmt.Println(string(newNiHao2[:2]))

	//UTF-8是一种变长编码，采用1-4个字节来存储一个字符，比如 𠮷 字
	//UTF-32 就属于定长编码，即永远用 4 字节存储码点，而 UTF-8、UTF-16 就属于变长存储，UTF-8 根据不同的情况使用 1-4 字节，而 UTF-16 使用 2 或 4 字节来存储码点。
	//定长和变长  "断句问题" 00000000111111110000101000011

	/*
	   [104 101 108 108 111]
	   [104 101 108 108 111]
	   [20320 22909]
	   [228 189 160 229 165 189]
	   he
	   你好
	   丑
	*/

}
