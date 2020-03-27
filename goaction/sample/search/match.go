package search

import (
	"log"
)

//Result 保存搜索的结果
type Result struct {
	Field   string
	Content string
}

//matcher定义了要实现的
//新搜索类型的行为


//interface  接口 简述
//interface 是一种类型，更准确的说 interface 是一种具有一组方法的类型，这些方法定义了 interface 的行为。
//go 允许不带任何方法的 interface ，这种类型的 interface 叫 empty interface。 空接口
//如果一个类型实现了一个 interface 中所有方法，我们说 该类型实现了该 interface，
//所以所有类型都实现了 empty interface，因为任何一种类型至少实现了 0 个方法。
// go 没有显式的关键字用来实现 interface，只需要实现 interface 包含的方法即可  就是不需要像php等一样 在类上implements xxx 这种
// interface 变量存储的是实现者的值

//判断 interface 变量存储的是哪种类型
//一个 interface 被多种类型实现时，有时候我们需要区分 interface 的变量 究竟存储哪种类型的值，go 可以使用 comma,
// ok 的形式做区分 value, ok := em.(T)： em 是 interface 类型的变量，T代表要断言的类型，value 是 interface 变量存储的值，

if t, ok := i.(*S); ok {
fmt.Println("s implements I", t)
}
ok 是 true 表明 i 存储的是 *S 类型的值，false 则不是，这种区分能力叫 Type assertions (类型断言)。

如果需要区分多种类型，可以使用 switch 断言，更简单直接，这种断言方式只能在 switch 语句中使用。

switch t := i.(type) {
case *S:
fmt.Println("i store *S", t)
case *R:
fmt.Println("i store *R", t)
}

// ok 是 bool 类型表示是否为该断言的类型 T。

//声明了一个接口
//这个接口声明了结构 类型或者具名类型需要实现的行为
//声明接口，如果接口类型只包含一个方法，那么这个类型的名字以 er 结尾。 如matcher 中的er

type Matcher interface {
	//这个方法输入一个指向 Feed 类型值的指针 和 一个 string 类型的搜索项。这个方法返回两个值：一个指向 Result 类型值 的指针的切片，另一个是错误值。
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 函数，为每个数据源单独启动 goroutine 来执行这个函数
//并发 的执行搜索
//注意 该函数 的第一个参数 是一个接口类型
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	//对特定的匹配器执行搜索
	//注意该接口类型 可以调用search方法
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}
	//将结果写入通道
	for _, result := range searchResults {
		results <- result
	}
}

// Display 从每个单独的 goroutine 接收到结果后 在终端窗口输出
func Display(results chan *Result) {
	// 通道会一直被阻塞，直到有结果写入 // 一旦通道被关闭，for 循环就会终止
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
