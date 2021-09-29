package main

import "fmt"

func main() {
	//双重循环
	fmt.Println("请输入三角形的高：")
	var count int
	fmt.Scanln(&count)
	for i := 1; i <= count; i++ {
		for j := 1; j <= 2*i-1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
