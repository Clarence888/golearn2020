package main

import "fmt"

//理解slice原理
//我们此处定义一个appendInt 函数 类比 原生的append函数 理解其内部操作
//注意：我么不清楚一次append会不会导致新的内存分配 也不能假设新老slice是不是指向同一底层数组
//因此：我们通常将append的调用结果再次赋值给传入append函数的slice 像这样 runes = append(runes,r)
//任何有可能改变slice的长度或者容量 都需要更新slice变量

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d len=%d , cap=%d\t%v\n", i, len(y), cap(y), y)

		x = y
	}
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
