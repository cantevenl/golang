package main

import "fmt"

func test1(n1 int) int {
	n1 = 30
	fmt.Println("test---", n1)
	return n1
}

func main() {
	n1 := 10
	n2 := test1(n1)

	fmt.Println("main里的n1是：", n1, "n2是：", n2)
}
