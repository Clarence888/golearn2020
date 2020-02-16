package main

/*
 curl --data "users=yifan,yifan2,changliang" http://localhost:9999/import
 curl http://localhost:9999/lucky
*/
import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"sync"
	"time"
)

//用户列表
var userList []string

//对共享变量读写的时候，需要加锁 防止并发导致数据出错。 保持线程安全性。

//线程不安全
//加互斥锁
var mu sync.Mutex

type lotteryController struct {
	Ctx iris.Context
}

//功能

func newApp() *iris.Application {
	app := iris.New()
	//启动根目录 处理
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()
	//userList = make([]string,0)
	userList = []string{}
	//初始化 互斥锁
	mu = sync.Mutex{}

	//运行起来
	app.Run(iris.Addr(":9999"))
}

func (c *lotteryController) Get() string {
	count := len(userList) // 总抽奖人数
	return fmt.Sprintf("当前总共参与抽奖的用户数：%d\n", count)
}

//导入用户名单
//post http/xxxxxx/import
//params users
func (c *lotteryController) PostImport() string {
	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",")

	//加锁
	mu.Lock()
	defer mu.Unlock()
	count1 := len(userList)

	for _, u := range users {
		u = strings.TrimSpace(u)
		if len(u) > 0 {
			userList = append(userList, u)
		}
	}

	count2 := len(userList)
	return fmt.Sprintf("当前总共参与的用户数：%d,成功导入的用户数：%d\n", count1, count2)
}

/*
抽奖
*/
func (c *lotteryController) GetLucky() string {
	//加锁
	mu.Lock()
	defer mu.Unlock()

	count := len(userList)
	if count > 1 {
		seed := time.Now().UnixNano()
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count))
		user := userList[index]
		userList = append(userList[0:index], userList[index+1:]...)
		return fmt.Sprintf("当前中奖用户：%s,剩余用户数：%d\n", user, count-1)
	} else if count == 1 {
		user := userList[0]
		return fmt.Sprintf("当前中奖用户：%s,剩余用户数：%d\n", user, count-1)
	} else {
		return fmt.Sprintf("没用用户参与啦，可利用import导入")
	}
}
