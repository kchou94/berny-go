package main

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Say() string { return c.Name + "：喵喵喵" }

type Dog struct {
	Name string
}

func (d Dog) Say() string { return d.Name + "：汪汪汪" }

func main() {
	c := Cat{Name: "小花猫"}
	fmt.Println(c.Say())
	d := Dog{Name: "小黑狗"}
	fmt.Println(d.Say())
}
