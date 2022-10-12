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

	// 重置目录，不然没权限创建文件夹
	err = os.Chdir("/Users/kchou/Code/mac-golang/src/berny-go")
	if err != nil {
		fmt.Println(err)
	}

	// 创建文件夹
	err = os.Mkdir("test", 0755)
	if err != nil {
		fmt.Println(err)
	}

	// 删除文件夹或文件
	err = os.Remove("test")
	if err != nil {
		fmt.Println(err)
	}

	// 修改文件夹或文件的名称
	err = os.Mkdir("test", 0755)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Rename("test", "go-test")
	if err != nil {
		fmt.Println(err)
	}

	// 新建文件
	_, err = os.Create("go-test/test.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 打开文件并写入文件
	/*
		O_RDONLY 打开只读文件
		O_WRONLY 打开只写文件
		O_RDWR 打开既可以读取又可以写入文件
		O_APPEND 写入文件时将数据追加到文件尾部
		O_CREATE 如果文件不存在，则创建一个新的文件
	*/
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString("你好")
	if err != nil {
		fmt.Println(err)
	}

	// 清理
	err = os.Remove("go-test/test.txt")
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove("test")
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove("go-test")
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove("file.txt")
	if err != nil {
		fmt.Println(err)
	}
}
