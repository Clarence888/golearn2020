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

type Point struct {
	X, Y int
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

	//=============
	//结构体字面量 设置结构体类型的值
	//1.必须按照顺序赋值
	p := Point{1, 2}
	//2. 指定名称和值 如果没指定 就是 该类型的零值
	v := Point{X: 1, Y: 3}

	fmt.Println(p.Y)

	fmt.Println(v.Y)
	//如果结构体的所有成员变量可比较 那么这个结构体 可比较  ==  != 其中  == 按照顺序比较两个结构体的变量的成员变量
	fmt.Println(p == v) //false 当成员变量的类型和值都相同 那么他俩相同

	//可比较的类型 都可以作为map的键类型
	hits := make(map[Point]int)
	hits[Point{2, 3}]++

	//结构体嵌套
	/*
		type Circle struct {
			X,Y,Radius int
		}

		type Wheel struct {
			X,Y,Radius, Spokes int
		}
	*/

	//有重复。。重构相同部分

	type DingWei struct {
		X, Y int
	}
	/*
		type Circle struct {
			Center DingWei
			Radius int
		}

		type Wheel struct {
			CicleAA Circle
			Spokes int
		}

		var w Wheel
		w.CicleAA.Center.X = 8
		//访问有点麻烦
	*/

	//go支持定义不带名称的结构体成员 只需要指定类型
	//于是乎
	type Circle struct {
		DingWei //匿名成员 不一定是结构体类型 任何命名类型或者指向命名类型的指针都可以
		Radius  int
	}

	type Wheel struct {
		Circle
		Spokes int
	}

	var ww Wheel
	ww.X = 8 //可以直接访问 等同于 ww.Circle.X
	fmt.Println(ww.X)

}
