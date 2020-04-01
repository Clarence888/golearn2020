package main

//判断变量的类型

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	//<目标类型的值>，<布尔参数> := <表达式>.( 目标类型 ) // 安全类型断言
	//interface{}代表空接口，任何类型都是它的实现类型
	value, ok := interface{}(container).(map[string]string)
	//把container变量的值转换为空接口值的interface{}(container)
	//ok 类型判断的结果
	//如果是true，那么被判断的值将会被自动转换为[]string类型的值，并赋给变量value，否则value将被赋予nil（即“空”）。
	/*
		没有ok  会引发 恐慌
			value := interface{}(container).([]string)

	*/

	fmt.Println(ok)
	fmt.Println(value)

	fmt.Printf("The element is %q.\n", container[1])
}
