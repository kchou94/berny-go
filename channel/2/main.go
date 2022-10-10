package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40
	ch <- 50
	close(ch)
	for i := range ch { // channel关闭后，会自动退出for循环
		fmt.Println(i)
	}
}
