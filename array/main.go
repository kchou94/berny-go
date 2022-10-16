package main

import "fmt"

func main() {
	// 数组的定义：var 变量名 [元素数量]类型
	var a [5]int
	fmt.Println(a)
	fmt.Printf("%T \n", a)
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)

	// 声明无长度的数组
	//var b [...]int
	//b[0] = 1
	//fmt.Println(b)
	var b = [...]int{1, 2, 3}
	fmt.Println(b)

	// 数组是值类型, 数据传递时属于值传递
	var arr = [3]int{1, 2, 3}
	arr2 := arr
	arr2[0] = 5
	fmt.Println(arr, arr2)
}
