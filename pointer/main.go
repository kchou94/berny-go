package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 10
	fmt.Println(a, &a, *(&a), reflect.TypeOf(a), reflect.TypeOf(&a), reflect.TypeOf(*(&a))) // 10 0xc0000140a8 10 int *int int

	a1 := 10
	b := &a1
	fmt.Println(a1, &a1, b)
	fmt.Println(a1, &b)
}
