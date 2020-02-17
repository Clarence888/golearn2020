package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//获取一个url内容 并输出

func main() {
	for _, url := range os.Args[1:] {
		//http.get产生一个http请求 返回存储在resp里
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:reding %s : %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
		fmt.Printf("%s", resp.Status)
	}
}
