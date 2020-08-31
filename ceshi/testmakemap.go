package main

import (
	"fmt"
)

func main() {

	var aa map[string]bool
	fmt.Println(aa == nil)

	a := make(map[string]int)
	fmt.Println(a == nil)

	a["hello"] = 1
	fmt.Println(a)
	a = make(map[string]int)
	fmt.Println(a)
}
