package main

import (
	"fmt"

	"github.com/bernylinville/berny-go/oop/encapsulation"
)

func main() {
	s := encapsulation.NewStudent("张三", 90)
	fmt.Println(s)
}
