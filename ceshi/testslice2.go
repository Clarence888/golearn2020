package main

import "fmt"

func main()  {
	sss := []int{1,2,3,4,5}

	//fmt.Printf("%T\n",sss)
	//fmt.Println(&sss)
	for i,v := range sss {
		fmt.Println(&i)
		//fmt.Println(&v)
		fmt.Println(sss)
		if i == 0 {
			sss = sss[:3]
			sss[2] = 9999
		}
		fmt.Println(i,v)
		fmt.Println("first",sss)

	}
	fmt.Println("---",sss)
}
