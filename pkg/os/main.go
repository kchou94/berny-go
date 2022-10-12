package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取当前目录
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pwd)

	// 修改当前目录
	err = os.Chdir("/Users")
	if err != nil {
		fmt.Println(err)
	}
	pwd, err = os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pwd)

	// 创建文件夹
	err = os.Mkdir("test", 0755)
	if err != nil {
		fmt.Println(err)
	}
}
