package main

import "fmt"

/*
import (
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	fmt.Println("hello")

	// Println 写到标准日志记录器
	log.Println("message")
	// Fatalln 在调用 Println()之后会接着调用 os.Exit(1)
	log.Fatalln("fatal message")
	// Panicln 在调用 Println()之后会接着调用 panic()
	log.Panicln("panic message")
}

*/

func main() {
	A := 1
	foo := func() {
		A = 2  //将main函数的A变量 重新赋值
		A := 3 //声明一个新变量 A 虽然同名 赋值为3
		A = 4  // 将新声明的变量赋值为4
		fmt.Println(A)
	}
	foo()
	fmt.Println(A) //打印main函数中的变量A
}

// 4 2
