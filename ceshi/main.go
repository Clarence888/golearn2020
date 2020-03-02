package main

import (
	"fmt"
	"net/http"
)

func main() {
	//所有的url 都用handler 处理
	http.HandleFunc("/", handler)
	//注意打日志只是为了监测  这里启动了服务器监听端口 监听8000
	//listens on the TCP network address and then calls
	//Serve with handler to handle requests on incoming connections.
	//会把请求给到上面handler 的
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}
