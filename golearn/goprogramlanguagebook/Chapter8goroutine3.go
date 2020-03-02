package main

//chan  通道 是goroutine之间的通信机制
//通道也有类型

func main() {
	//创建一个通道 是个引用 复制和传参都是操作的引用 通道的零值是nil 同一种类型的通道可以比较 ==
	ch := make(chan int)
	//make  无缓冲通道  unbuffered
	//第二个参数 表述通道容量 缓冲通道
	//ch2 := make(chan int,10)

	//通信
	//send
	ch <- x //把x发送给通道
	//receive
	x = <-ch //赋值
	<-ch     //丢弃结果

	//close关闭通道
	close(ch)

}

//注意 发送操作 将会阻塞goroutine 直到被接收 值操作完成 各自goroutine继续执行
//无缓冲通道也叫做 同步通道
//每一个消息有一个值 但是有时候 没有额外信息的时候 单纯目的是同步 可以用一个 struct{} 类型的通道来强调

//通道用来连接 goroutine  一个的输出是另一个的输入  ---》管道
