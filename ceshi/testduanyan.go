package main

import "fmt"

func main() {
	var i interface{}
	i = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	i = 100
	t, ok := i.(int)
	fmt.Println(t, ok)

	/*
		一个 interface 可被多种类型实现，有时候我们需要区分 interface 变量究竟存储哪种类型的值？类型断言提供对接口值的基础具体值的访问
		t := i.(T)
		该语句断言接口值i保存的具体类型为T，并将T的基础值分配给变量t。
		如果i保存的值不是类型 T ，将会触发 panic 错误。为了避免 panic 错误发生，可以通过如下操作来进行断言检查

	*/
	t2 := i.(string) //panic
	fmt.Println(t2)
}
