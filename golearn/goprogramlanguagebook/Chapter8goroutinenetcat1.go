package main

import (
	"io"
	"log"
	"net"
	"os"
)

//自己实现一个netcat

func main() {
	conn, err := net.Dial("tcp", "localhost:9111")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

//当我们启动这个程序时候  启动两个客户端 但是只有一个可以输出时间

//因为服务器是顺序的 一次处理一个请求

//如何支持并发呢？在服务端程序做改造
