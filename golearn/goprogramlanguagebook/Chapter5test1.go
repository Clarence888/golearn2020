package main

import (
	"fmt"
	"math"
)

//函数
//方法 函数名前多了一个参数 这个参数把这个方法 绑定到这个参数对应的类型上
type Point struct {
	X, Y float64
}

//普通函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Point类型的  方法 p称为方法的接收者  //可以把他看作是 xxx类 定义的 里面的xxx方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 两个名字一样为何不冲突？
//第一个是 包级别的函数  名字可以认为是 包.Distance
//第二个声明一个Point类型的方法 真实名字 Point.Distance
func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	//p.Distance 称为 选择子  selector
}

//注意 每个类型都有自己的命名空间  所以方法的接受者不同 那么就不会冲突
//同一个类型拥有的方法名 必须要唯一
type QiTa struct {
	A, B int
}

func (q QiTa) Distance() int {
	return q.A
}

//同一个包下的任何类型都可以声明方法 只要他的类型不是 指针 或者 接口 类型 。

//注意 下面这种是不行的
//invalid receiver type 'Z' ('Z' is a pointer type)
type Z *int

//func (Z) ZhiZhen() int{
//
//}

//这种是可以的 因为 Point不是指针类型 这里只是为了避免复制实参 采用指针传递变量地址
func (p *Point) CeShi() int {
	p.X = 111
}

//整型链表
//*IntList 的类型nil 代表 空列表
type IntList struct {
	Value int
	Tail  *IntList
}

//
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
