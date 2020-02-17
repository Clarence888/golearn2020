package main

//扩充 maintest2 扩充不采用流式读取 直接读取整个输入到大块内存
//找出重复行
//输出标准输入中出现次数大于1 的行  前面是次数
import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//定义了一个键值对集合 键是 每行输入的东西 值是 出现的次数
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		//读取文件 整个内容
		//ReadFile函数返回一个(可以转化成字符串的) 字节 slice
		//ReadAll 类似方法
		//todo 延展学习：操作文件相关的使用File自带的Read方法 使用bufio库的Read方法 使用io/ioutil库的ReadAll() 使用io/ioutil库的ReadFile()
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		//计数大于1 的 输出出来 出现几次 以及对应的行内容 也就是map里面的键
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
