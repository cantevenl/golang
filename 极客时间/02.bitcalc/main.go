package main

import "fmt"

func main() {
	a, b := 9, 4
	fmt.Println(a ^ b)
	fmt.Println(b ^ a) //亦或满足交换律

	arr := []int{8, 6, 3, 4, 3, 4, 5, 5, 6}
	result := -1
	for _, item := range arr {
		if result < 0 {
			result = item
			fmt.Println(result)
		} else {
			result = result ^ item
			fmt.Println(result)
		}
	}
	fmt.Println(result)
}
