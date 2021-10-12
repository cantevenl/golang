package main

import "fmt"

func printPyramid(high int) {
	fmt.Println("输入金字塔的高：")
	fmt.Scanln(&high)

	for i := 1; i <= high; i++ {
		for k := 1; k <= high-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == high {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	printPyramid(10)
}
