package main

import "fmt"

func main()  {
	/*
	var age int = 1;
	for i := 1; i<10;i++ {
		age = age +2
		ab := 1
		fmt.Println(ab)
	}
	fmt.Println(age)
	fmt.Println(ab) //作用域不对
	fmt.Println(i) //作用域不对

	*/
	var p *int
	v := 2
	p = &v
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Printf("%p \n",p)
	fmt.Printf("%#p \n",p)

	var m map[string]int
	m = make(map[string]int)
	m["aa"] = 1
	m["bb"] = 2
	mp := &m
	fmt.Println(m)
	fmt.Println(mp)
	fmt.Println(*mp)
	fmt.Printf("%p \n",mp)
	fmt.Printf("%#p \n",mp)
}
