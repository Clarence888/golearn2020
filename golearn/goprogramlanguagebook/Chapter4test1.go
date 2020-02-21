package main

import "fmt"

//复合数据类型 数组 slice map 结构体
//数组 //固定长度 拥有零个或者多个 相同数据类型的元素的序列
//len 可返回数组中 元素个数

func main() {
	var a [3]int             //3个整数的数组
	fmt.Println(a[0])        //输出第一个元素
	fmt.Println(a[len(a)-1]) //输出最后一个元素

	fmt.Printf("%d \n", &a[0])

	//输出索引和元素
	//对于所有的 range 循环，Go 语言都会在编译期间将原切片或者数组赋值给一个新的变量 ha，在赋值的过程中其实就发生了拷贝，所以我们遍历的切片其实已经不是原有的切片变量了。
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
		fmt.Printf("%d  %d\n", &i, &v)
	}

	//仅仅输出元素
	for _, v := range a {
		fmt.Printf("%d,\n", v)
	}
	/*
		初始的都是0值

		0
		0
		824634327072
		0 0
		824634318880  824634318888
		1 0
		824634318880  824634318888
		2 0
		824634318880  824634318888
		0,
		0,

	*/

	//var aa [3]int = [3]int{1,2,3}
	var bb [3]int = [3]int{1, 2}
	fmt.Println(bb[1])
	fmt.Println(bb[2]) // 0

	var cc = [...]int{1, 2, 4, 5, 6} //... 代表 数组长度由初始的后面大括号的元素个数决定
	fmt.Println(cc[2])

	var dd = [...]string{99: "hello"}

	fmt.Println(dd[99])
	fmt.Println(dd[98])

	ca := [2]int{1, 2}
	cd := [3]int{1, 2, 3}

	fmt.Println(ca == cd) // invalid operation: ca == cd (mismatched types [2]int and [3]int)

	//数组传递到函数 传值 传的是副本

}
