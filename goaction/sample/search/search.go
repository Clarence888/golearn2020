package search

import (
	"log"
	"sync"
)

//注册用于搜索的匹配器映射
//包级变量 Mather类型的映射 键类型为string
//标识符 首字母大学 公开 能被其他包代码直接访问  小写 不公开
var matchers = make(map[string]Matcher)

//执行搜索逻辑
func Run(searchTerm string) {

	//获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}


	//创建一个无缓冲通道 接受匹配后的结果
	results := make(chan *Result)

	//构造一个waitgroup 以便处理所有的数据源
	//计数信号量
	var waitGroup sync.WaitGroup

	//设置需要等待处理
	//每个数据源的goroutine数量
	waitGroup.Add(len(feeds))

	//为每个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		//获取一个匹配器用于查找
		//matchers 返回map命中的值 exists 返回是否存在在map里
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		//启动一个goroutine 执行搜索
		go func(matcher Matcher, feed *Feed) {
			//这里有用到闭包 因为闭包 函数可以直接访问那些没有作为参数传入的变量
			Match(matcher, feed, searchTerm, results)
			//递减waitGrou计数
			waitGroup.Done()

		}(matcher, feed)
	}

	//启动一个goroutine 来监控是否所有工作都做完了
	go func() {
		//等待所有任务完成
		//wait方法会导致goroutine阻塞 直到 waitgroup内部的计数到达 0
		waitGroup.Wait()
		//用关闭通道的方式 通知下面的Display函数 //可以退出程序
		close(results)
	}()

	//启动函数 显示返回结果，并且在最后一个结果显示完 返回

	Display(results)
}

// Register 调用时，会注册一个匹配器，提供给后面的程序使用

//将一个 Matcher 值加入到保存注册匹配器的映射中。
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
