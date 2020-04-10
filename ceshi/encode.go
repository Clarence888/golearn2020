package main

import (
	"encoding/json"
	"fmt"
)

type PersonInfo struct {
	ID      string
	name    string
	Address string
}

func main() {
	myMap := make(map[string]int)
	myMap["One"] = 1
	myMap["two"] = 2

	fmt.Printf("%#v\n", myMap)

	encoded, _ := json.Marshal(myMap)
	fmt.Println(string(encoded)) // {"One":1}    // 私有字段 two 被忽略了

	//json.Unmarshal(encoded, &out)
	//fmt.Printf("%#v\n", out)     // main.MyData{One:1, two:""}
}
