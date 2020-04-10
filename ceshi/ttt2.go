package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"time"
)

type Student struct {
	Name string
	Age  int
}

var cls chan Student = make(chan Student, 2)

func main() {

	//var aaaChan chan int
	//aaaChan = make(chan int)
	//aaaChan <- 2
	////fmt.Println(aaaChan)

	var data []int = []int{1, 2, 3, 4, 5, 6, 7}

	var intChan chan int = make(chan int)

	stringChan := make(chan string)

	stu := new(Student)
	stu.Age = 21
	stu.Name = "Andy"
	fmt.Println(stu)
	fmt.Printf("%v\n", stu)
	fmt.Printf("%+v\n", stu)
	fmt.Printf("%p \n", stu)
	aa := 3
	fmt.Printf("%T\n", aa)
	zhizhena := &aa
	fmt.Printf("%T\n", zhizhena)
	benlai := &zhizhena
	fmt.Println(benlai)
	fmt.Printf("%T\n", benlai)

	go func() {
		aa := stu
		fmt.Printf("发送type:%T \n", aa) //
		cls <- *stu
		stu.Age = 20
		fmt.Println("----", stu.Age)
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ss := <-cls
		fmt.Printf("%T \n", ss)
		fmt.Println(ss)

	}()
	time.Sleep(10 * time.Second)
	fmt.Println("最后的type", stu.Age)
}
