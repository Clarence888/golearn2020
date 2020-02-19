package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

/*
//实现一个miniweb服务器
//http://localhost:8000/hello
//go run xxxx.go & 可选择在后台执行
func main()  {
	//所有的url 都用handler 处理
	http.HandleFunc("/",handler)
	//注意打日志只是为了监测  这里启动了服务器监听端口 监听8000
	//listens on the TCP network address and then calls
	//Serve with handler to handle requests on incoming connections.
	//会把请求给到上面handler 的
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"URL.path = %q\n",r.URL.Path)
}

*/

//扩展一下 加一个计数器
var mu sync.Mutex

//共享变量  可能会产生并发访问
var count int

func main() {
	//所有的url 都用handler 处理
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	//注意打日志只是为了监测  这里启动了服务器监听端口 监听8000
	//listens on the TCP network address and then calls
	//Serve with handler to handle requests on incoming connections.
	//会把请求给到上面handler 的
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//处理响应 次数是返回请求的路径信息
//todo 参数是什么意思呢
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)

	//r.Method r.URL r.Proto r.Header 可能常用
}

//回显目前为止的调用次数
func counter(w http.ResponseWriter, r *http.Request) {
	//此处为互斥锁 开箱即用 只需要声明 var mu sync.Mutex
	//todo 深入学习总结 互斥锁 另外还有读写锁var rwm = sync.RWMutex
	//rwm.Lock()
	//rwm.Unlock()
	//rwm.RLock()
	//rwm.RUnlock()
	mu.Lock()
	fmt.Fprintf(w, "count %d\n", count)
	mu.Unlock()
	//此处unlock也可以用
	//defer mu.Unlock()
	//	fmt.Println("还没结束")
}

/*
这些程序中,我们看到了作为输出流的三种非常不同的类型。fetch程序复制HTTP响
应到文件os. Stdout,像lissajous一样; fetchall程序通过将响应复制到iout1. Discard中
进行丢弃(在统计其长度时):Web服务器使用fmt.Fprintf通过写入httpResponsewriter来
让浏览器显示
*/
/*
handler:=func(w http.Responsewriter, r *http.Request) {
lissajous(w)
}
http.Handlefunc("/",handler)

或者
http.Handlefunc("/",func(w http.Responsewriter,r *http.Request){
lissajous (w)
})
*/
