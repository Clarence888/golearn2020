package main

import "fmt"

func main() {
	//var data []int = []int{1,2,3,4,5,6,7}
	//tmpData := data
	//sli1 := data[:3]
	//sli2 := tmpData[4:]
	//
	//fmt.Println(sli1)
	//fmt.Println(sli2)
	//
	//newData := append(sli1,sli2...)
	//fmt.Println(newData)
	//
	//var data []int = []int{1,2,3,4,5,6,7}
	//fmt.Println(len(data), cap(data), data)
	//
	//data[2],data[6] = data[6],data[2]
	//data = data[0:6]
	//fmt.Println(len(data), cap(data), data)

	//s1 := data[:3]
	//fmt.Println(len(s1), cap(s1), s1)
	//
	//s2 := data[4:]
	//fmt.Println(len(s2), cap(s2), s2)

	//data = append(data[:2],data[3:]...)

	//fmt.Println(data)
	/*
		names := [4]string{
			"John",
			"Paul",
			"George",
			"Ringo",
		}
		fmt.Println(names)
		a := names[0:1]
		fmt.Println(len(names))
		b := names[1:2]
		fmt.Println(a, b)


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

	*/

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
