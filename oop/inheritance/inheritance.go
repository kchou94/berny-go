package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func (Person) SayHello() {
	fmt.Println("This is from Person")
}

type Student struct {
	Person
	school string
}

func main() {
	stu := Student{school: "清华大学"}
	stu.name = "张三"
	stu.age = 18
	stu.sex = "男"
	fmt.Println(stu)

	stu.SayHello()
}
