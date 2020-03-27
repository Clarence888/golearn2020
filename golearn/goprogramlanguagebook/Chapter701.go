package main

import (
	"fmt"
	"io"
	"xorm.io/builder"
)

/*
接口

go的接口是隐式实现的

对于一个具体的类型 无需声明它实现了哪些接口 只要提供接口所必须实现的方法即可，

接口类型是一种 抽象类型
一个接口类型定义了一套放他  如果一个具体的类型要实现该接口 那么必须实现接口类型的所有方法

*/

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Coser interface {
	Close() error
}

//组合现有接口 嵌入式接口
type ReadAndCloser interface {
	Reader
	Closer
}

//如果一个类型 实现了接口要求的所有方法 那么这个类型 实现了这个接口

//接口赋值规则 仅仅当一个表达式实现了一个接口时 这个表达式才可以赋值给该接口

/*
Go不是一种典型的OO语言，它在语法上不支持类和继承的概念。
没有继承是否就无法拥有多态行为了呢？答案是否定的，Go语言引入了一种新类型—Interface，它在效果上实现了类似于C++的“多态”概念，虽然与C++的多态在语法上并非完全对等，但至少在最终实现的效果上，它有多态的影子。
虽然Go语言没有类的概念，但它支持的数据类型可以定义对应的method(s)。本质上说，所谓的method(s)其实就是函数，只不过与普通函数相比，这类函数是作用在某个数据类型上的，所以在函数签名中，会有个receiver(接收器)来表明当前定义的函数会作用在该receiver上。
Go语言支持的除Interface类型外的任何其它数据类型都可以定义其method（而并非只有struct才支持method），只不过实际项目中，method(s)多定义在struct上而已。
从这一点来看，我们可以把Go中的struct看作是不支持继承行为的轻量级的“类”。

Interface定义了一个或一组method(s)，这些method(s)只有函数签名，没有具体的实现代码（有没有联想起C++中的虚函数？）。
若某个数据类型实现了Interface中定义的那些被称为"methods"的函数，则称这些数据类型实现（implement）了interface。
这是我们常用的OO方式，
*/

//demo
//1
type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func(s S) Get()int {
	return s.Age
}

func(s *S) Set(age int) {
	s.Age = age
}

//3
func f(i I){
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	s := S{}
	f(&s)  //4
}

//#1 定义了 interface I，在 #2 用 struct S 实现了 I 定义的两个方法，
//接着在 #3 定义了一个函数 f 参数类型是 I，S 实现了 I 的两个方法就说 S 是 I 的实现者，执行 f(&s) 就完了一次 interface 类型的使用。

s := S{}
var i I //声明 i
i = &s //赋值 s 到 i
fmt.Println(i.Get())


//interface 的变量中存储的是实现了 interface 的类型的对象值，这种能力是 duck typing。
//在使用 interface 时不需要显式在 struct 上声明要实现哪个 interface ，只需要实现对应 interface 中的方法即可，go 会自动进行 interface 的检查，并在运行时执行从其他类型到 interface 的自动转换，即使实现了多个 interface，go 也会在使用对应 interface 时实现自动转换，这就是 interface 的魔力所在。



//demo2
type Animal interface {
	Bark() string
	Walk() string
}

type Dog struct {
	name string
}

//Dog 实现了 Animal 接口，所以可以用 Animal 的实例去接收 Dog的实例，必须是同时实现 Bark() 和Walk() 方法，否则都不能算实现了Animal接口。

func (dog Dog) Bark() string {
	fmt.Println(dog.name + ":wan wan wan!")
	return "wan wan wan"
}

func (dog Dog) Walk() string {
	fmt.Println(dog.name + ":walk to park!")
	return "walk to park"
}

func main() {
	//声明一个 Animal 类型的 变量
	var animal Animal

	//只声明没赋值的interface 是nil interface，value和 type 都是 nil
	//只要赋值了，即使赋了一个值为nil类型，也不再是nil interface
	fmt.Println("animal value is:", animal)	//animal value is: <nil>
	fmt.Printf("animal type is: %T\n", animal) //animal type is: <nil>

	//该变量接受dog数据类型赋值
	animal = Dog{"旺财"}
	animal.Bark() //旺财:wan wan wan!
	animal.Walk() //旺财:walk to park!

	fmt.Println("animal value is:", animal) //animal value is: {旺财}
	fmt.Printf("animal type is: %T\n", animal) //animal type is: main.Dog
}



//demo3
nil interface
在上面的例子中，我们打印刚定义的 animal:

value为 nil
type 也为 nil

官方定义：Interface values with nil underlying values:

只声明没赋值的interface 是nil interface，value和 type 都是 nil
只要赋值了，即使赋了一个值为nil类型，也不再是nil interface

type I interface {
	Hello()
}

type S []int

func (i S) Hello() {
	fmt.Println("hello")
}
func main() {
	var i I
	fmt.Printf("1:i Type:%T\n", i)
	fmt.Printf("2:i Value:%v\n", i)

	var s S
	if s == nil {
		fmt.Printf("3:s Value%v\n", s)
		fmt.Printf("4:s Type is %T\n", s)
	}

	i = s
	if i == nil {
		fmt.Println("5:i is nil")
	} else {
		fmt.Printf("6:i Type:%T\n", i)
		fmt.Printf("7:i Value:%v\n", i)
	}
}

output:
1:i Type:<nil>
2:i Value:<nil>
3:s Value[]
4:s Type is main.S
6:i Type:main.S
7:i Value:[]
复制代码从结果看，初始化的变量 i 是一个 nil interface,当把值为 nil 的变量 s 赋值i后,i 不再为nil interface。
细心的同学，会发现一个细节，输出的第3行
3:s Value[]
复制代码明明，s的值是 nil,却输出的是一个[],这是由于 fmt使用反射来确定打印的内容，因为 s 的类型是slice，所以 fmt用 []来表示。

