package main

import (
	"html/template"
	"os"
	"strings"
)

//注意 这里的title 和 .Name1 要空开 否则识别失败
//模板定义。双层大括号 点xxx  模板渲染  比如 .Name1 .HELLO
//函数调用funcMap 注册自定义函数 title
const templateText = `
Output 0: {{title .Name1}}
Output 1: {{title .Name2}}
Output 2: {{.Name3 | title}}
`

func main() {
	funcMap := template.FuncMap{"title": strings.Title}
	//strings.Title函数 首字母大写的副本
	tpl, _ := template.New("go-programming-tour").Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "go",
		"Name2": "programming",
		"Name3": "tour",
	}
	_ = tpl.Execute(os.Stdout, data)

}
