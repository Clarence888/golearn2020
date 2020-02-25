package main

import (
	"fmt"
	"sort"
)

//map 散列表 键值对元素 无序集合 注意：键唯一
//go 语言中 map散列表的引用 map[k]v  k:键的类型 v：值的类型  k的类型必须是可以通过 == 来比较的类型
//map 类型的零值是 nil 获取元素个数 len
func main() {

	//ages := make(map[string]int)  //创建一个从string 映射 int的map
	//ages["alice"] = 21
	//ages["charlie"] = 34
	ages := map[string]int{
		"alice":   21,
		"charlie": 34,
		"harry":   44,
		"adm":     14,
	}

	//新的空map  map[string]int{}

	//访问
	ages["charlie"] = 32
	fmt.Println(ages["charlie"])

	//删除一个元素
	//delete(ages,"charlie")
	//fmt.Println(ages["charlie"]) //0 被取值的时候也可以不存在（这样就返回对应类型的零值

	//map元素不是一个变量 不可以获取地址
	//_ = &ages["bob"] //编译错误
	//为什么呢？因为map的增长可能会导致 已有元素被重新散列到其他新位置

	for a, b := range ages {
		fmt.Println(a, b)
	}
	//注意 通过range遍历map时，key的顺序被随机化
	//业务依赖key次序时，如何解决随机化问题?
	//引入slice 作为key的顺序列
	sorted_keys := make([]string, 0)
	for k, _ := range ages {
		sorted_keys = append(sorted_keys, k)
	}

	// sort 'string' key in increasing order
	sort.Strings(sorted_keys)

	for _, k := range sorted_keys {
		fmt.Printf("k=%v, v=%v\n", k, ages[k])
	}

}

/*
为什么这很特别？

一句话：态度。

这个无伤大雅的语言特性在我看来恰是作为通用语言哲学的一个闪光点。没有过于灵活地允许马马虎虎的代码，Go强迫你从一开始就把事情弄得直接一点。Go程序员参考里面说如果他们的程序可以编译（而且代码符合Go的风格），那么代码有很大的可能可以像预期的那样工作，这种模模糊糊却不错的感觉也有Go严谨性的贡献。没有诡异的类型bug，丢失分号等等错误。

尤其是，在Andrew的参考文章中，他指出这是Go的设计者们所作出的改变。他们不再允许人们依赖于那些破破烂烂的假设。我最痛恨的一点就是那些破烂的，到处是bug的功能（这发生在交付的产品中，或者是编程语言，等等很多地方），通过权衡，接受，从而变成了一个特性，另外尝试修复这些"特性"真的是很恶心。很明显的，PHP和JavaScript的语言文化就是因为各种原因往这个方向发展的（他们使用它，但是注定要付出代价，而且很多东西到最后都是没有解决的）。

例如，PHP的一个最大的缺点是，针与干草堆的问题（在干草堆里面找针）。我理想中的语言所应该具有的特点和这种不一致性格格不入。这也是为什么我发现Go的设计者拒绝糟糕的异常和泛型设计。他们就是想做正确的事情，当然他们知道这需要花费时间。他们不着急，而且向语言中添加特性比删除特性容易多了。

hash map 类似c++里的unordered_map，是无序的，开发者不希望有人觉得有序是内置map的特性
*/

/*
makemap 这个函数返回的结果：*hmap，它是一个指针，而我们之前讲过的 makeslice 函数返回的是 Slice 结构体：
func makeslice(et *_type, len, cap int) slice
复制代码回顾一下 slice 的结构体定义：
// runtime/slice.go
type slice struct {
    array unsafe.Pointer // 元素指针
    len   int // 长度
    cap   int // 容量
}
复制代码结构体内部包含底层的数据指针。
makemap 和 makeslice 的区别，带来一个不同点：当 map 和 slice 作为函数参数时，在函数参数内部对 map 的操作会影响 map 自身；而对 slice 却不会（之前讲 slice 的文章里有讲过）。
主要原因：一个是指针（*hmap），一个是结构体（slice）。Go 语言中的函数传参都是值传递，在函数内部，参数会被 copy 到本地。*hmap指针 copy 完之后，仍然指向同一个 map，因此函数内部对 map 的操作会影响实参。而 slice 被 copy 后，会成为一个新的 slice，对它进行的操作不会影响到实参。

*/
/*map 不可比较 只可以和 nil比较  判断两个map是不是有相同的键值  需要些循环*/
/*map   if age,ok := ages["bob"];!ok {xxxxxx}
判断bob在不在map中
* /
*/
