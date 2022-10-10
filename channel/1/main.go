package main

import "fmt"

func main() {
	// 1. 创建channel
	ch := make(chan int, 5)
	// 2. 向channel放入数据
	ch <- 10
	ch <- 20
	fmt.Println("发送成功", ch)
	// 3. 从channel取出数据
	v1 := <-ch
	fmt.Println("v1=", v1)
	v2 := <-ch
	fmt.Println("v2=", v2)
	// 4. 空channel取值报错
	v3 := <-ch
	fmt.Println("v3=", v3)
}
