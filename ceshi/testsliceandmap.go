package main

import "fmt"

func main() {
	var s []int
	s = append(s, 1)

	var m map[string]int
	//m = make(map[string]int)
	m["one"] = 1 //错误
	fmt.Println(m["one"])
}
