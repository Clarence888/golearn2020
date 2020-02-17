package main

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
	//bufio可以简单高效处理输入输出
	input := bufio.NewScanner(os.Stdin)

	//输入的
	for input.Scan() {
		// # 作结束符
		if strings.TrimSpace(input.Text()) == "#" || input.Err() != nil {
			break
		}
		//如果没有此作为结束符 可ctrl+d
		counts[input.Text()]++
	}

	//注意 忽略怒普通.err()中可能的错误

	for line, n := range counts {
		//计数大于1 的 输出出来 出现几次 以及对应的行内容 也就是map里面的键
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
