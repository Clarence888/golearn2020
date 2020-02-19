package main

import "fmt"

//指针

func main() {
	x := 1
	p := &x         //注意 此处声明p 是整型指针 指向x 这个变量（也就是p包含x的地址）
	fmt.Println(*p) //输出 1
	*p = 2          //将p指向的地址 赋值为2
	fmt.Println(x)
}
