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
	input := bufio.NewScanner(os.Stdin) //此处处理标准输入

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

	//for 使用方法 todo  range todo
	//range 内置函数 遍历数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。  类似php的foreach？
	//当用于遍历数组和切片的时候，range函数返回索引和元素；
	//当用于遍历字典的时候，range函数返回字典的键和值。
	// range函数用来遍历字符串时，返回Unicode代码点。
	// 第一个返回值是每个字符的起始字节的索引，第二个是字符代码点，
	// 因为Go的字符串是由字节组成的，多个字节组成一个rune类型字符。
	//for index, value := range mySlice
	//for…range中那个value，是一个值拷贝，而不是元素本身。

	for line, n := range counts {
		//此处遍历的是字典 n是值 line是键
		//计数大于1 的 输出出来 出现几次 以及对应的行内容 (也就是map里面的键)
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
