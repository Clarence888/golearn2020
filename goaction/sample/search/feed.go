package search

import (
	"encoding/json"
	"os"
)

//const dataFile = "data/data.json"
const dataFile = "/Users/clarence/go/src/goaction/sample/data/data.json"

//Feed包含我们的需要处理的数据源信息
//会解码到feed结构组成的切片里
//定义结构类型 注意大写字母 会对外暴露
// `` 里面的部分被称作 标记 tag
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

//读取并 反序列化 元数据文件
//无参数 两个返回值 第一个是一个切片 每一项指向一个Feed类型的值  error用来表示函数是否调用成功
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)

	if err != nil {
		return nil, err
	}

	//在函数执行返回前首先检查当前函数内是否有推迟执行的函数，如果有则执行，然后再返回
	//可以很方便的在函数结束前执行一些清理工作，比如关闭打开的文件、关闭连接、释放资源、解锁等等。
	//一个函数中可以使用多个defer，其执行顺序与栈类似：后进先出 //即便是意外崩溃了 也会执行
	//可读性高  可在适当位置书写 不用像传统一样等到最后才定义相关函数
	//关闭文件
	defer file.Close()

	//将文件解码到一个切片里
	//这个切片的每一项是一个指向一个Feed类型的指针
	var feeds []*Feed
	//返回值上调用 Decode方法
	err = json.NewDecoder(file).Decode(&feeds)

	//此函数无须检查错误 调用方会检查

	return feeds, err
}
