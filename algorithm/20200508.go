package main

import "fmt"

//字符串反转

func main()  {

	str := "abcd"

	for k,v := range str {
		fmt.Println("k:",k,"v:",v)
	}
	str[3] = 100

	fmt.Println(str)
}