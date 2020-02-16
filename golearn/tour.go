package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"time"
)

//https://tour.go-zh.org/basics

//import可以用圆括号 也可以 分开每行导入
//import "fmt"
//import "math"

//声明变量 或者变量列表 最后是类型
//未明确初始值的变量 会被赋予 "零值" 比如 int 0 string "" 布尔 false
var c, python, java, php bool
var i int

//声明可以包含初始值 有初始值的可以不再写类型
var aa, bb, cc int = 1, 2, 3
var dd = "nihao"

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//常量声明 不可用:=
const Pi = 3.14

//数值常量  高精度的值
const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

type Vertex struct {
	X int
	Y int
}

func main() {

	var ct int

	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	//在 Go 中，如果一个名字以大写字母开头，那么它就是已导出的 在导入一个包时，你只能引用其中已导出的名字。任何“未导出”的名字在该包外均无法访问。
	fmt.Println(math.Pi)
	fmt.Println(add(11, 13))
	fmt.Println(test())
	a, b := swap("hellonihao", "chifanlema")
	fmt.Println(a, b)
	fmt.Println(split(27))
	fmt.Println(i, c, python, java, php, ct, aa, bb, cc, dd)
	test2()

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	changetest()

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	xunhuan()

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	testSwitch()

	//结构体 结构体的字段用点访问
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	//结构体字段可以通过结构体指针来访问。
	p := &v
	p.X = 1e9
	fmt.Println(v)

	qiepian()

	testMap()

	changeMap()
}

//:=  短变量声明 代替var 该结构不能在函数外使用
//函数可以没有参数或接受多个参数。
func test() string {
	mm := "duanbianliang"
	return mm
}

//类型在变量名 之后
//如果都相同 可缩写 x, y int
func add(x int, y int) int {
	return x + y
}

func add1(x, y int) int {
	return x + y
}

//函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}

//返回值可被命名 如下的x y  应当具有一定的意义 可被使用
//没有参数的 return 语句返回已命名的返回值。也就是 直接 返回 会返回 x y
//应当仅用在下面这样的短函数中 否则可能会影响可读性
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//短变量声明
func test2() {
	c, python, java := true, false, "no!"
	fmt.Println(c, python, java)
}

/*
go 基本类型
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
    // 表示一个 Unicode 码点

float32 float64

complex64 complex128
*/

//类型转换
func changetest() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	/*或者这样
	i := 42
	f := float64(i)
	u := uint(f)*/

	//类型推导
	//var i int
	//j := i // j 也是一个 int

	fmt.Println(x, y, z)
}

//循环 go 只有for循环
func xunhuan() {
	sum := 0
	//无需小括号 if也是 无需小括号
	for i := 0; i < 100000; i++ {
		sum += i
		//if 不需要小括号
		if i < 10 {
			println(i)
		}

		//if 简短赋值 可以在条件表达式前执行一个简单的语句  作用于仅仅在if {} 和 else {}内
		//if v := 33; v < 199 {
		//	println("helloooooooo");
		//}
	}

	/*可选
		sum := 1
			for ; sum < 1000; {
				sum += sum
			}
	//或者这样
		for sum < 1000 {
			sum += sum
		}

		如果直接写for 就相当于 while
		for{
		//死循环
		}
	*/
	fmt.Println(sum)
}

//if else
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

//switch Go 的 switch 的 case 无需为常量，且取值不必为整数。
//Go 只运行选定的 case，而非之后所有的 case
//Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止

func testSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	//没有条件的 switch 同 switch true 一样。
	//这种形式能将一长串 if-then-else 写得更加清晰。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}

//Go 语言的 defer 会在当前函数或者方法返回之前执行传入的函数。它会经常被用于关闭文件描述符、关闭数据库连接以及解锁资源
//defer代码块的作用域仍然在函数之内 defer 传入的函数不是在退出代码块的作用域时执行的，它只会在当前函数和方法返回之前被调用。
//defer 可认为是在return之后执行 或者是 外层函数返回前
//1.调用 defer 关键字会立刻对函数中引用的外部参数进行拷贝
/*
func a() {
i := 0
defer fmt.Println(i) //输出0，因为i此时就是0
i++
defer fmt.Println(i) //输出1，因为i此时就是1
return
}*/
//2.defer执行顺序为先进后出 循环中多次调用 倒序 4 3 2 1 //先进后出 栈
/*
func b() {
for i := 0; i < 4; i++ {
defer fmt.Print(i)
}
}*/
//3.defer可以读取有名返回值
/*
func c() (i int) {
defer func() { i++ }()
return 1
}*/

func testDefer() {

}

//go 指针

//类型 *T 是指向 T 类型值的指针。其零值为 nil。
//var p *int
/*
& 操作符会生成一个指向其操作数的指针。

i := 42
p = &i
* 操作符表示指针指向的底层值。

fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
这也就是通常所说的“间接引用”或“重定向”。
*/

//Go 没有指针运算。
func testPointer() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

//一个结构体（struct）就是一组字段（field）。
/*
type Vertex struct {
	X int
	Y int
}
*/

//数组
//[n]T 表示拥有 n 个 T 类型的值的数组
func testArr() {
	//拥有几个string类型的值的数组
	//数组不能改变大小
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//数组直接赋值
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

//切片
//类型 []T 表示一个元素类型为 T 的切片
/*
切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：

a[low : high]
它会选择一个半开区间，包括第一个元素，但排除最后一个元素。

以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：

a[1:4]
*/

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

/*
切片文法
切片文法类似于没有长度的数组文法。

这是一个数组文法：

[3]bool{true, true, false}
下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：

[]bool{true, true, false}*/

//切片就像数组的引用  并不存储任何数据 更改切片的元素会修改其底层数组中对应的元素
//将切片理解成一片连续的内存空间加上长度与容量的标识。
//容量当做成总长度减去左指针走过的元素值
//扩容不可超过底层数组的容量 panic: runtime error: slice bounds out of range [:19] with capacity 6
func qiepian() {
	////切片的零值是 nil nil 切片的长度和容量为 0 且没有底层数组
	//var s []int

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var sm []int = primes[1:4]
	fmt.Println(sm)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	//切片下界的默认值为 0，上界则是该切片的长度。
	/*
		对于数组

		var a [10]int
		来说，以下切片是等价的：

		a[0:10]
		a[:10]
		a[0:]
		a[:]
	*/
	//切片长度 包含的元素的个数
	//容量 第一个元素开始数，到其底层数组元素末尾的个数
	//切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取

	sa := []int{2, 3, 5, 7, 11, 13}
	printSlice(sa)

	// 截取切片使其长度为 0
	sa = sa[:0]
	printSlice(sa)

	// 拓展其长度
	sa = sa[:4]
	printSlice(sa)

	// 舍弃前两个值
	sa = sa[2:]
	printSlice(sa)

	//超过扩容最大值
	//s = s[2:100]
	//printSlice(s)

}

//用 make 创建切片
//a := make([]int, 0,4)
//make([]类型，长度，容量)

//切片的切片

//向切片追加元素 追加后结果是一个包含原切片所有元素加上新添加元素的切片
//当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组
//func append(s []T, vs ...T) []T
//元素类型为T的切片
//s = append(s, 0,1)

//遍历切片 for 循环的range形式
func testRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	//_ 来忽略
	//for i, _ := range pow
	//for _, value := range pow

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}

//映射 map 零值为 nil 。nil 映射既没有键，也不能添加键

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func testMap() {
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	//映射的文法与结构体相似，不过必须有键名
	var m2 = map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		"Google": Vertex2{
			37.42202, -122.08408,
		},
	}
	//类型名 顶级类型是一个类型名 可省略
	//"Bell Labs": {40.68433, -74.39967},
	//	"Google":    {37.42202, -122.08408},
	fmt.Println(m2)
}

//映射修改

func changeMap() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	//通过双赋值检测某个键是否存在：
	//
	//elem, ok = m[key]
	//若 key 在 m 中，ok 为 true ；否则，ok 为 false。
	//
	//若 key 不在映射中，那么 elem 是该映射元素类型的零值。
	//
	//同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

//函数也是值。它们可以像其它值一样传递。
//函数值可以用作函数的参数或返回值。
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func testFunc() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

//函数的闭包
//Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。
//
//例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。
/*
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

*/
