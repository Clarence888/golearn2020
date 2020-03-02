package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	//counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals) //循环结束后关闭通道 会导致squares结束循环 并引发关闭squares通道 最终 主goroutine结束
	}()

	//squares
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	//printer
	for x := range squares {
		fmt.Println(x)
	}
}
