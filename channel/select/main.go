package main

import "fmt"

func main() {
	// 在某些场景下我们需要同时从多个通道接收数据,这个时候就可以用到golang中给我们提供的select多路复用

	// 1. 定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 2. 定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道时,如果不关闭会阻塞而导致deadlock
	// 问题: 在实际开发中,可能我们不好确定什么时候关闭该管道
	// 解决方法: 可以使用select方式可以解决
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%v\n", v)
		default:
			fmt.Println("数据获取完毕")
			return // 如果不加return,会一直循环
		}
	}
}
