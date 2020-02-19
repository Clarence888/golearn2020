package main

//flag 包相关 todo 需要学习
import (
	"flag"
	"fmt"
	"strings"
)

//创建了一个布尔类型的标识变量  三个参数 名字 默认值 介绍
var n = flag.Bool("n", false, "omit trailing newline")

//创建了一个字符串变量
var sep = flag.String("s", "---", "separator")

//flag.xxxx() The return value is the address of a xxxx variable
//注意 变量 sep  n 是指向标识变量的指针 如果要访问他们 需要 *sep *n

func main() {
	//更新标识变量的默认值
	flag.Parse()
	//strings.join 连接字符串  第二个为分隔符
	//Join 将第一个参数中的子串连接成一个单独的字符串，子串之间用 sep 分隔
	//注意 这里 的*sep 就是上面定义的  s 参数对应的值 是--- 所以
	fmt.Print(strings.Join(flag.Args(), *sep))

	//*n 默认是false 此处的判断是 如果是true 就输出
	if !*n {
		fmt.Println()
	}
}

/*
❯ go run maintest9.go -s / a bc def 注意 这是-s 是指用什么分割 后面的/ 是自定义的分隔符
a/bc/def

~/go/src/golearn/goprogramlanguagebook master*
❯ go run maintest9.go -n a bc def
a-- bc-- def

~/go/src/golearn/goprogramlanguagebook master*
❯ go run maintest9.go -help
Usage of /var/folders/l5/cjpkf3g55xzdfqmspnht_wd00000gn/T/go-build846752459/b001/exe/maintest9:
  -n    omit trailing newline
  -s string
        separator (default " ")
exit status 2

*/
