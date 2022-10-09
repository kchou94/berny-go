package main

import "fmt"

func main() {
	var a interface{} = "hello"
	v, ok := a.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("非字符串类型")
	}
}
