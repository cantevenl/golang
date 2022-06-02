package main

import "fmt"

const MAX int = 3

func main() {
	var a int = 20
	var p *int
	p = &a
	fmt.Printf("%T,%v,%v", p, p, *p)

	var arr = []int{1, 2, 3}
	var ptr [MAX]*int
	fmt.Println(ptr)
	for i := 0; i < MAX; i++ {
		ptr[i] = &arr[i]
	}

	for i := 0; i < len(ptr); i++ {
		fmt.Println(*ptr[i])
	}
}
