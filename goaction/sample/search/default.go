package search

//实现了默认匹配器
//使用一个空结构声明了一个名叫 defaultMatcher 的结构类型
//空结构 在创建实例时，不会分配任何内存。
type defaultMatcher struct {
}

//init 函数将默认匹配器注册到程序里
//程序里所有的 init 方法都会在 main函数启动前被执行
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search 实现了默认匹配器的行为
//defaultMatcher 类型实现 Matcher 接口的代码
//Search 方法的声明也声明了 defaultMatcher类型的值 的接收者
//如果声明函数的时候带有接收者，(此处为(m defaultMatcher))，则意味着声明了一个方法。
//这个方法会和指定的接收者的 类型绑在一起。
//在我们的例子里， Search 方法与 defaultMatcher 类型的值绑在一起。
//这意 味着我们可以使用 defaultMatcher 类型的值或者指向这个类型值的指针来  调用 Search 方 法。
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}

//大部分方法在被调用后都需要维护接收者的值的状态，所以，一个最佳实践是，将方法 的接收者声明为指针。

//go语言中 函数 方法 是不同的
//方法的声明和函数类似，他们的区别是：方法在定义的时候，会在func和方法名之间 增加一个参数，这个参数就是 [接收者]，
//这样我们定义的这个方法就和接收者绑定在了一起，称之为 这个接收者的方法。
/*
type person struct {
	name string
}

func (p person) Show() string{
	return "the person name is "+p.name
}
留意例子中，func和方法名之间增加的参数(p person),这个就是 接收者。现在我们说，类型person 有了一个Show方法，现在我们看下如何使用它。

func main() {
	p:=person{name:"张三"}
	fmt.Println(p.Show())
}
调用的方法非常简单，使用类型的变量进行调用即可，类型变量和方法之前是一个.操作符，表示要调用这个类型变量的某个方法的意思
*/

/*
Go 语言中同时有函数和方法。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。语法格式如下：

func (variable_name variable_data_type) function_name() [return_type]{
   //函数体
}
下面定义一个结构体类型和该类型的一个方法：
package main

import (
"fmt"
)

type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 = ", c1.getArea())
}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}
*/
