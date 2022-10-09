package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	name   string // 私有属性不能被 json 包访问
	Sno    string
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "小王子",
		Sno:    "S001",
	}
	fmt.Printf("%#v\n", s1)

	s, err := json.Marshal(s1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v\n", err)
		return
	}
	jsonStr := string(s)
	fmt.Println(jsonStr)
}
