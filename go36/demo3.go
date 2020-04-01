package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var err error
	n, err := io.WriteString(os.Stdout, "Hello, everyone!\n")
	//对err的重声明
	fmt.Println(n, err)

	var name string

	name = "changliang"

	name := "xiaoer"
	fmt.Println(name)

}
