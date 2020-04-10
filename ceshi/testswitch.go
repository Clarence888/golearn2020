package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	//os作用域只在switch里
	//fmt.Printf("%s.---1", os)

	switch os := runtime.GOOS; {
	case "darwin" == os:
		fmt.Println("OS X.")
	case "linux" == os:
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	//os作用域只在switch里
	//fmt.Printf("%s.----2", os)

}
