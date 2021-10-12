package main

import "fmt"

func sum(n1, n2 int) int {
	defer fmt.Println("defer n1=", n1)
	defer fmt.Println("defer n2=", n2)
	n1++
	n2++
	res := n1 + n2
	fmt.Println("res=", res)
	return res
}

func main() {
	res := sum(27, 9)
	fmt.Println(res)
}
