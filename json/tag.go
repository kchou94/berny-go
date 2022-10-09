package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int    `json:"id"` // 通过指定 tag 实现 json 序列化该字段时的 key
	Gender string `json:"gender"`
	Name   string
	Sno    string
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		Name:   "小王子",
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
