package main

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() i = ", i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go test() // 开启一个协程
	for i := 0; i < 2; i++ {
		fmt.Println("main() i = ", i)
		time.Sleep(time.Millisecond * 100)
	}
}
