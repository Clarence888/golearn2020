package main

import (
	"fmt"
	"time"
)

//struct 结构体
//零个或者多个任意类型的命名变量 组合在一起 的 聚合数据类型
//每个变量叫做 结构体成员

//demo
//通常一个成员变量 占据一行
type Employee struct {
	ID        int
	Name      string
	Address   string //首字母大写 可导出
	Dob       time.Time
	Position  string
	Salary    int
	ManagerID int
}

//命名

type ZhuGuan struct {
	Id   int
	Name string
	age  int
	//zg ZhuGuan //注意 类型内部不可以包含他本类型 如果想实现递归 可以用指针
	zhizhenzhuguan *ZhuGuan
}

func main() {
	//实例化
	var dilbert Employee
	//如何访问 点号
	dilbert.Address = "北京市海淀区"
	dilbert.Salary = 100000
	//获取变量地址
	xinshui := &dilbert.Salary
	fmt.Println(dilbert.Address)
	fmt.Println(dilbert.Salary)
	fmt.Println(&dilbert.Salary)
	*xinshui = 2000000 //加薪
	fmt.Println(dilbert.Salary)

	//点号用在结构体指针
	//这里我们把 dilbert 赋值给了employeeOfTheMonth  它是个指针
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position = "nihaowa"
	fmt.Println(dilbert.Position)

	//--------------
	var dazhuguan ZhuGuan
	dazhuguan.Name = "niendadad"
	fmt.Println(dazhuguan.Name)
	dazhuguan.zhizhenzhuguan = &dazhuguan
	dazhuguan.zhizhenzhuguan.Id = 1111
	fmt.Println(dazhuguan.Id)
}
