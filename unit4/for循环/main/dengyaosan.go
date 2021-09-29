package main

import "fmt"

func main() {
	fmt.Println("请输入三角形的高：")
	var high int
	fmt.Scanln(&high)
	for i := 1; i <= high; i++ {
		for j := 1; j <= high-i; j++ {
			fmt.Printf(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
