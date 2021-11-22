package main

import "fmt"

type MethodUtils struct {
	num1 int
	num2 int
}

func (jiu MethodUtils) Jiu(num int) {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %v \t", j, i, j*i)
		}
		fmt.Println()
	}
}

func main() {
	var jiu MethodUtils
	jiu.Jiu(9)
}
