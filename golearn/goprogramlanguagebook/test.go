package main

import "fmt"

type Tt struct {
	Age int
}
type BuBian struct {
}

func main() {
	//obj := new(Tt)
	//obj2 := new(Tt)
	//
	//bubian1 := new(BuBian)
	//bubian2 := new(BuBian)
	//fmt.Println(&obj)
	//fmt.Println(&obj2)
	//fmt.Println("fenge-----")
	//
	//fmt.Println(bubian1)
	//fmt.Println(bubian2)
	//time.Sleep(time.Second*10)

	//s := "helloworld"
	//fmt.Println(s[0],s[4])
	for i, r := range "Hello,世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
	/*
	      0       'H'     72
	      1       'e'     101
	      2       'l'     108
	      3       'l'     108
	      4       'o'     111
	      5       ','     44
	      6       '世'    19990
	      9       '界'    30028

	   //一个汉字占3字节（byte）
	*/
	//
}
