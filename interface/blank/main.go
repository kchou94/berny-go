package main

import "fmt"

// 空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	show(20)
	show("hello")
	slice := []int{1, 2, 3}
	show(slice)

	// 切片实现空接口
	var slice1 = []interface{}{1, "hello", true, 20.5}
	fmt.Println(slice1)

	// 空接口作为map的值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "小明"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}
