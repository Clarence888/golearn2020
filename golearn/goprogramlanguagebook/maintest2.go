package main

//扩充 maintest1 扩充可以从文件读出
//找出重复行
//输出标准输入中出现次数大于1 的行  前面是次数
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//定义了一个键值对集合 键是 每行输入的东西 值是 出现的次数
	counts := make(map[string]int)
	//定义文件
	//注意： os.Args
	//os.Args 提供原始命令行参数访问功能。 os.Args的类型是 []string ，也就是字符串切片
	//注意，切片中的第一个参数是该程序的路径，也就是文件本身的路径。
	//一般使用 os.Args[1:]保存所有程序的的参数。
	//如何使用  os.Args[3] 类似这样既可 用标准的索引位置方式
	//说白了 类似于 PHP脚本里面 argv[0]  这种
	//取除了执行文件本身以外的所有参数
	files := os.Args[1:]

	if len(files) == 0 {
		//如果没有参数 也就是没有赋值  那么就取标准输入
		countLines(os.Stdin, counts)
	} else {
		//有参数
		//流式 模式读取输入
		for _, arg := range files {
			//os.Open 返回两个值  一个是文件本身的引用 一个是 内置的error类型值
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		//计数大于1 的 输出出来 出现几次 以及对应的行内容 也就是map里面的键
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	/*
		for idx, args := range os.Args {
			fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
		}
	*/
}

//计算每行的数据 出现了多少字
//注意思考 这里 map被传送给一个函数了，函数接收到的是这个引用的副本，
//那么 对他的改变  main函数也是可见的吗(函数内部改了 为何在外部也变了)？
//作用域问题？内存？
//为什么？todo
func countLines(f *os.File, counts map[string]int) {

	input := bufio.NewScanner(f)
	//输入的
	for input.Scan() {
		// # 作结束符
		if strings.TrimSpace(input.Text()) == "#" || input.Err() != nil {
			break
		}
		//如果没有此作为结束符 可ctrl+d
		counts[input.Text()]++
	}
}
