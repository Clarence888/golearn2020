package main

import "fmt"

func main() {
	//查看数组最后一个非0元素的位置
	//遍历数组 发现为0 将末尾非0元素赋值给此位置 并向前挪动一位非零元素
	//直到重合
	var testSlice []int

	testSlice = []int{0, 3, 9, 8, 0, 7, 6, 0, 4, 0, 2, 1}
	fmt.Println(testSlice)

	j := 0
	for k, v := range testSlice {
		if v != 0 {
			testSlice[j] = v
			if k != j {
				testSlice[k] = 0
			}
			j++
		}
	}

	fmt.Println(testSlice)
}
