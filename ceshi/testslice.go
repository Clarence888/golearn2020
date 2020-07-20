package main

import "fmt"

func main()  {
	var arr [5]int = [5]int{1,2,3,4,5}

	s1 :=  arr[:]
	s2 := arr[1:4]

	sfroms1 := s1[1:4]



	fmt.Println("arr is ",arr)
	fmt.Println("s1 is ",s1)
	fmt.Println("s2 is ",s2)
	fmt.Println("sfroms1 is ",sfroms1)

	//arr[2] = 222
	//
	//fmt.Println("arr is ",arr)
	//fmt.Println("s1 is ",s1)
	//fmt.Println("s2 is ",s2)
	//
	//
	//s1[2] = 333
	//s2[1] = 444

	sfroms1[2] = 999
	fmt.Println("arr is ",arr)
	fmt.Println("s1 is ",s1)
	fmt.Println("s2 is ",s2)



}
