package main

import "fmt"

//slice 切片 拥有相同类型元素的 可变长度 的序列
//[]T 比如 []string []int 和数组区别是 []里面不需要写信息 也就是长度等

//有个slice底层数组

//slice 指针 长度 容量  len cap 返回

//一个底层数组可以对应多个slice

//slice操作符 s[i:j] 创建一个新slice

//slice 唯一可以允许做比较 是 nil    oneslice == nil  因为其初始值是nil
//值为nil的slice没有底层数组。 长度和容量都是零

//检查slice是否为空  用  len(s) == 0

//make  make([]T,len) //容量可以省略 省略情况下 容量和长度相同 make([]T,2,3)
//也就是make创建了一个无名数组并返回了他的一个slice

func main() {
	//months := [...]string{1:"january",2:"febary",12:"December"}
	//初始化slice s 注意 [] 中没有长度
	s := []int{0, 1, 8, 3, 4, 5}
	fmt.Println(s[2])

	ss := make([]int, 3)
	ss[1] = 2
	fmt.Println(ss[0], ss[1])

}
