package main

import (
	"fmt"

	"berny-go/oop/encapsulation"
)

func main() {
	s := encapsulation.NewStudent("张三", 90)
	fmt.Println(s)
	fmt.Println(s.GetScore())
	s.SetScore(80)
	fmt.Println(s)
}
