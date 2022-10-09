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
	jsonStr := `{"id":1,"gender":"男","Name":"张三","Sno":"S001"}`
	var s Student
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
	}
	// s.Name = "李四"
	fmt.Printf("%#v \n", s)
}
