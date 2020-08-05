package main

import "fmt"

func main() {
	var english = "hello"
	fmt.Println("英文的hello长度是:", len(english))

	var englishAndChinese = "hello 你好"
	fmt.Println("中英文混合的hello长度是:", len(englishAndChinese))

}
