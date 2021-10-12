package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Printf("hello world")
	fmt.Println(runtime.Version())

	a := 1
	fmt.Println(a, &a)
	a = 2
	fmt.Println(a, &a)
	b := &a
	*b = 10
	fmt.Println(a)
}
