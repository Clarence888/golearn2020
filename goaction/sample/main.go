package main

import (
	"goaction/sample/search"
	"log"
	"os"
)

//init 在main前调用
func init() {
	log.SetOutput(os.Stdout)
}

//整个程序的入口

func main() {
	//search.Run("president")
	search.Run("More")
}
