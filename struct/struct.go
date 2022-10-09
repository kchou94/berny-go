package main

import "fmt"

type person struct {
	name string
	city string
	age  int
}

func main() {
	var p1 person
	p1.name = "小王子"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%+v\n", p1)
	fmt.Printf("p1=%#v\n", p1)
	fmt.Println("--------------------")
	p2 := new(person)
	p2.name = "小王子"
	p2.city = "北京"
	p2.age = 18
	fmt.Printf("p2=%v\n", p2)
	fmt.Printf("p2=%+v\n", p2)
	fmt.Printf("p2=%#v\n", p2)
	fmt.Println("--------------------")
}
