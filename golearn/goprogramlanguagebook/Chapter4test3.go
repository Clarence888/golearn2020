package main

import (
	"fmt"
)

//理解slice原理
//我们此处定义一个appendInt 函数 类比 原生的append函数 理解其内部操作
//注意：我么不清楚一次append会不会导致新的内存分配 也不能假设新老slice是不是指向同一底层数组
//因此：我们通常将append的调用结果再次赋值给传入append函数的slice 像这样 runes = append(runes,r)
//任何有可能改变slice的长度或者容量 都需要更新slice变量

//todo slice如何初始化？集中方法

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d len=%d , cap=%d\t%v\n", i, len(y), cap(y), y)

		x = y
	}

	var a = []int{1, 2, 3, 4, 5, 6, 7}
	var b = []int{1, 2, 3, 4, 5, 6, 7}
	a = removeEasy(a, 3)
	b = removeBaoShunXu(b, 3)
	fmt.Println(a) //[1 2 3 7 5 6]
	fmt.Println(b) //[1 2 3 5 6 7]

}

func appendInt(x []int, y int) []int {
	//入参 源slice  增加的int值
	//定义一个切片
	var z []int

	//我们这里是在算 如果我这个切片追价一个int 长度是什么情况
	zlen := len(x) + 1

	//%p是打印地址的, %x是以十六进制形式打印, 完全不同！另外在64位下结果会不一样, 所以打印指针老老实实用%p
	fmt.Printf("打印切片的地址 %p \n", z)

	if zlen < cap(x) {
		//如果当前切片的长度加1 之后 比源slice的容量小 说明这时候 还能再次扩充此slice
		//size仍有增长空间 扩充slice内容
		z = x[:zlen]
		fmt.Printf("打印切片的地址 %p \n", z)

	} else {
		//如果追价一个后 长度等于源slice的容量或者大于 那么就得新分配空间
		//无空间了 为它分配一个新的底层数组
		//简单策略 为了达到分摊线性复杂度 容量扩展一倍
		//把现在的长度复制容量  如果两倍长度 能放下 新的容量就是两倍长度
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}

		//重新制作一个临时切片
		z = make([]int, zlen, zcap)

		fmt.Printf("扩充切片后 %p \n", z)
		//参数：目标切片 源切片 会将源切片的slice 复制到目标slice中
		copy(z, x) //内置copy函数
		//有返回值 返回实际复制的元素的个数 这个值是两个slice长度的最小值
	}
	z[len(x)] = y
	return z
}

//slice可以用来实现栈
//func stackDemo()  {
//	stack = make([]int,10,10)
//
//	stack = append(stack,v)  //相当于push v
//	top := stack[len(stack) - 1] //栈顶
//	//弹出最后一个  缩减栈
//	stack = stack[:len(stack)-1] //pop
//
//}

//如何移除一个slice中的某个元素
//1.并且保持剩余元素顺序
//利用copy 将高位向前移动
func removeBaoShunXu(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])

	return slice[:len(slice)-1]
}

//2.不保持剩余元素顺序
//直接将最后一个赋值给对应的索引
func removeEasy(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

/*
定义切片
你可以声明一个未指定大小的数组来定义切片：

var identifier []type
切片不需要说明长度。

或使用make()函数来创建切片:

var slice1 []type = make([]type, len)

也可以简写为

slice1 := make([]type, len)
也可以指定容量，其中capacity为可选参数。

make([]T, length, capacity)
这里 len 是数组的长度并且也是切片的初始长度。

切片初始化
s :=[] int {1,2,3 }
直接初始化切片，[]表示是切片类型，{1,2,3}初始化值依次是1,2,3.其cap=len=3

s := arr[:]
初始化切片s,是数组arr的引用

s := arr[startIndex:endIndex]
将arr中从下标startIndex到endIndex-1 下的元素创建为一个新的切片

s := arr[startIndex:]
默认 endIndex 时将表示一直到arr的最后一个元素

s := arr[:endIndex]
默认 startIndex 时将表示从arr的第一个元素开始

s1 := s[startIndex:endIndex]
通过切片s初始化切片s1

s :=make([]int,len,cap)
通过内置函数make()初始化切片s,[]int 标识为其元素类型为int的切片

len() 和 cap() 函数
切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。
*/
