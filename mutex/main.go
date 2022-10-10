package main

import (
	"fmt"
	"sync"
)

var num int

var mtx sync.Mutex
var wg sync.WaitGroup

func add() {
	defer wg.Done()
	// 互斥锁
	// 如果不加锁，会出现多个协程同时访问num变量，导致结果不正确
	mtx.Lock()
	num++
	mtx.Unlock()
}

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println("num = ", num)
}
