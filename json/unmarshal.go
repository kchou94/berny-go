package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	Name   string
	Sno    string
}

func main() {
	jsonStr := `{"ID":1,"Gender":"男","Name":"张三","Sno":"S001"}`
	var s Student
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v\n", err)
	}
	// s.Name = "李四"
	fmt.Printf("反序列化后 student=%#v student.Name=%v \n", s, s.Name)
}
