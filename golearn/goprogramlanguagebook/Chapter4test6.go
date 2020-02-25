package main

import (
	"fmt"
)

//二叉树实现插入排序
type tree struct {
	value       int
	left, right *tree
}

//结构体的零值 由 结构体成员的零值构成
//没有任何成员变量的结构体成为 空结构体  seen :=make(map(string)struct{})  只有键有用  这样吧map当做集合用

func main() {
	//初始化一个slice
	//var sa = []int{111,222,44,55,6,77777,1131,1223}
	var sa = []int{27, 3, 5, 11, 98}
	Sort(sa)        //排序
	fmt.Println(sa) //打印
}

//原地排序就是指不申请多余的空间来进行的排序，就是在原来的排序数据中比较和交换的排序。例如堆排序等都是原地排序，合并排序（根据TAOCP，合并排序也有原地排序的版本），计数排序等不是原地排序。
//属于原地排序的是：希尔排序、冒泡排序、插入排序、选择排序、堆排序、快速排序。
//原地排序
func Sort(values []int) {
	//声明一个树结构体
	var root *tree
	//把slice的值 全部赋值给这棵树
	for _, v := range values {
		root = add(root, v)
		fmt.Println(root.value) //注意 在本例子中 此处 root.value 一直为27

	}
	//清空切片 values[:0]
	appendValues(values[:0], root)
}

//将元素按照顺序追加到values里面 结果返回slice
func appendValues(values []int, t *tree) []int {
	//
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

//参数一 树  参数二 值 返回树
func add(t *tree, value int) *tree {
	//是空树
	if t == nil {
		//等价于返回 &tree{value:value}
		t = new(tree)
		//赋值第一个值
		t.value = value
		return t
	}

	//不是空树
	//当前值小于t的value 就赋给左节点
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		//当前值大于t的value 就赋给右节点
		t.right = add(t.right, value)
	}
	return t
}

//树的形状
/*
       27
     3      98
   0    5
      0    11
*/
