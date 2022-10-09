package main

import "fmt"

type Usb interface {
	Start()
	Stop()
	SetName(name string)
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name + "手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

func (p *Phone) SetName(name string) {
	p.Name = name
}

func main() {
	phone1 := Phone{Name: "小米"}
	var p1 Usb = &phone1
	p1.Start()
	p1.Stop()
	p1.SetName("华为")
	fmt.Println(phone1.Name)
}
