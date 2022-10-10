package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 第一步：声明一个计数器

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1() i = ", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // 第三步：计数器减一
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2() i = ", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // 计数器减一
}

func main() {
	wg.Add(1) // 第二步：计数器加一
	go test1()
	wg.Add(1) // 计数器加一
	go test2()
	wg.Wait() // 第四步：等待协程结束
	fmt.Println("主线程退出...")
}
