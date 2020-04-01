package main

import (
	"flag"
	"fmt"
	"go36/lib"
	"os"
)

func init() {
	flag.StringVar(&Name, "name", "everyone", "输入要打招呼的人名")
}

var Name string

func main() {

	//自定义命令源码文件的参数使用说明
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "这是可以自定义的help话术 Usage of %s:\n", "question")
		flag.PrintDefaults()
	}

	//解析参数
	flag.Parse()
	//fmt.Printf("Hello,%s\n",name)
	lib.Hello(Name)
	//如果把代码分到其他的代码文件里面 在go run 或者  go build 的时候 需要多个文件一起
	//go build demo.go demo_lib.go
}

/*
❯ go run demo.go -name=changliang
Hello,changliang

go run demo.go --help
Usage of /var/folders/l5/cjpkf3g55xzdfqmspnht_wd00000gn/T/go-build252379234/b001/exe/demo:   //这个是临时的可执行文件
-name string
输入要打招呼的人名 (default "everyone")
exit status 2

go build demo.go 会生成可执行文件

./demo -name=hhh
Hello,hhh
*/
