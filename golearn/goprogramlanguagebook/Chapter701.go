package main

import (
	"io"
	"xorm.io/builder"
)

/*
接口

go的接口是隐式实现的

对于一个具体的类型 无需声明它实现了哪些接口 只要提供接口所必须实现的方法即可，

接口类型是一种 抽象类型
一个接口类型定义了一套放他  如果一个具体的类型要实现该接口 那么必须实现接口类型的所有方法

*/

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Coser interface {
	Close() error
}

//组合现有接口 嵌入式接口
type ReadAndCloser interface {
	Reader
	Closer
}

//如果一个类型 实现了接口要求的所有方法 那么这个类型 实现了这个接口

//接口赋值规则 仅仅当一个表达式实现了一个接口时 这个表达式才可以赋值给该接口
