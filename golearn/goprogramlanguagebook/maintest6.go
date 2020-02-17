package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//并发访问获取多个url内容

func main() {
	start := time.Now()
	//创建字符串通道
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		//传进来几个网址 就开启几个goroutine
		go fetch(url, ch) //启动一个goroutine
	}
	//接受发送到通道的消息 并输出
	for range os.Args[1:] {
		fmt.Println(<-ch) //从通道ch接收
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //发送到通道ch
		return
	}
	//todo io.Copy 读取响应内容 写入discard输出流丢弃
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	resp.Body.Close() //不要泄露资源
	if err != nil {
		ch <- fmt.Sprintf("while reding %s:%v", url, err) //发送到通道ch
		return
	}
	secs := time.Since(start).Seconds()
	//每次访问成功 就发送一行汇总信息到通道ch
	ch <- fmt.Sprintf("%.2fs %7d %s ", secs, nbytes, url)
}
