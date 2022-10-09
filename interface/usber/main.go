package main

import "fmt"

// 1.接口是一个规范
type Usber interface {
	getName() string
}

// 2.如果接口里面有方法的话，必要要通过结构体或者通过自定义类型实现这个接口
type Phone struct {
	Name string
}

// 3.手机要实现usb接口的话必须得实现usb接口中的所有方法
func (p Phone) getName() string {
	return p.Name
}

type Computer struct {
	Brand string
}

func (c Computer) getName() string {
	return c.Brand
}

func main() {
	p := &Phone{Name: "小米"}
	c := &Computer{Brand: "联想"}

	var p1 Usber = p // go 中接口是一个数据类型
	fmt.Println(p1.getName())

	// 接口使用场景，处理相同类型的数据
	newName := transData(p)
	newName1 := transData(c)
	fmt.Println(newName, newName1)
}

func transData(usber Usber) string {
	name := usber.getName()
	return fmt.Sprintf("%s%s", name, "处理后")
}
