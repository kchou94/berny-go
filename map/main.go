package main

import "fmt"

func main() {
	// map 的定义 var 变量名 = map[key类型]value类型
	var userInfo = map[string]string{
		"username": "zhangsan",
		"password": "123456",
	}
	userInfo["gender"] = "Male"
	fmt.Println(userInfo)

	// 优先用 make 初始化
	var userInfo2 = make(map[string]string)
	userInfo2["username"] = "lisi"
	userInfo2["password"] = "123456"
	fmt.Println(userInfo2)

	// 判断某个 key 是否存在
	//v, ok := userInfo2["username"]
	v, ok := userInfo2["gender"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("key 不存在")
	}

	// 删除 map 中的元素，一组 key value
	delete(userInfo, "gender")
	fmt.Println(userInfo)

	// 遍历
	for key, value := range userInfo {
		fmt.Println(key, value)
	}
	for key := range userInfo {
		fmt.Println(key)
	}
	for _, value := range userInfo {
		fmt.Println(value)
	}
}
