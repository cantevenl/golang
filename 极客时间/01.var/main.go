package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := "hello"
	fmt.Println(reflect.TypeOf(a1))
	a2 := 3
	fmt.Println(reflect.TypeOf(a2))
	a3 := 3.0
	fmt.Println(reflect.TypeOf(a3))
	a4 := &a3
	fmt.Println(reflect.TypeOf(a4))
}
