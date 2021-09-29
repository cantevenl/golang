package main

import "fmt"

func main() {
	var num int = 10
	fmt.Println("num的值是：", num)
	change(&num)
	fmt.Println("num的值是：", num)
}

func change(n *int) {
	*n = 30
}
