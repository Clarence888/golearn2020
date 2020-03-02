package main

import (
	"io"
	"log"
	"net"
	"time"
)

//go 支持两种并发编程风格
//Go 语言提供了 sync 和 channel 两种方式支持协程(goroutine)的并发。
//go里 每一个并发执行的活动称之为goroutine

//demo1 顺序时钟服务器

func main() {
	//创建一个对象 在一个端口上监听进来的连接
	listener, err := net.Listen("tcp", "localhost:9111")
	if err != nil {
		log.Fatal(err)
	}
	for {
		//accept被阻塞 直到有请求进来 就返回出去一个对象 来代表一个连接
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		//不支持并发
		//handleConn(conn)
		//支持并发
		go handleConn(conn)
	}
}

//处理一个完整的连接
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		//format参数是一个模板
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

/*
netcat  也可以自己实现一个netcat 用net.Dial
❯ nc localhost 9111
13:45:45
13:45:46
13:45:47
13:45:48
13:45:49
13:45:50
13:45:51
13:45:52
13:45:5
*/
