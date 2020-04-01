package main

import "fmt"

/*
变量的声明周期：通过他是否能可达来确定
包级别的变量----》整个程序的执行时间
局部变量有一个动态的生命周期:每次执行声明语句时创建一个新的实体,
变量一直生存到它变得不可访问,这时它占用的存储空间被回收。
函数的参数和返回值是局部变量,它们在其闭包函数被调用的时候创建。

垃圾回收器怎么知道一个变量应该被回收？

//todo 内存管理 关键词
//todo 逃逸 涉及堆栈存储
func f(){
var x int
x=1
golbal - &x  我们说 x从f中逃逸了
}
*/

//类型声明 通常出现在包级别
//type 名字 类型

//虽然都是float64 但是不是相同的类型 所以不能用算术表达式比较以及合并
//命名类型的底层类型决定了它的结构和表达方式,以及它支持的内部操作集合 比如下面的底层类型是float64

type Celsius float64    //摄氏温度
type Fahrenheit float64 //华氏温度

const (
	AbsoluteZeroC Celsius = -273.5
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	//对于每个类型T,都有一个对应的类型转换操作T(x)将值x转换为类型T
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	//对于没有
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	var c Celsius
	var f Fahrenheit
	//命名类型的值可以和 与他相同类型的值  或者 底层类型相同的未命名类型的值 比较
	//不同命名类型的值 不能直接比较
	fmt.Println(c == 0) //true
	fmt.Println(f >= 0) //true
	//fmt.Println(c == f)  //编译错误 类型不匹配
	fmt.Println(c == Celsius(f)) //true
}
