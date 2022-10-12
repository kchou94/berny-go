package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// 读取文件中所有数据
	fileName1 := "./test.txt"
	data, err := ioutil.ReadFile(fileName1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	// 写入数据
	s1 := "Hello world!"
	err = ioutil.WriteFile(fileName1, []byte(s1), 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	// 读取所有数据
	s2 := "qwertyuiop"
	r1 := strings.NewReader(s2)
	data2, err := ioutil.ReadAll(r1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data2)

	// 读取一个目录下子内容：子文件和子目录，但是仅有一层
	dirName := "/Users/kchou"
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(fileInfos))
}
