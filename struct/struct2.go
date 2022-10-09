package main

import "fmt"

type Person struct {
	name string
	age  int
}

// 值类型接收者
func (p Person) printInfo() {
	fmt.Printf("姓名:%v 年龄:%v\n", p.name, p.age)
}

// 指针类型接收者
func (p *Person) setInfo(name string, age int) {
	p.name = name
	p.age = age
}

func main() {
	p1 := Person{"小王子", 18}
	p1.printInfo()
	p1.setInfo("张三", 20)
	p1.printInfo()
}
