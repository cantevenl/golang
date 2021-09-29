package main

import "fmt"

func main() {
	//第1种
	var arr1 [3]int = [3]int{2, 4, 6}
	fmt.Println(arr1)
	//第2种
	var arr2 = [3]int{3, 6, 9}
	fmt.Println(arr2)
	//第3种
	var arr3 = [...]int{4, 5, 6, 7}
	fmt.Println(arr3)
	//第4种
	var arr4 = [...]int{2: 66, 0: 33, 1: 99, 3: 88}
	fmt.Println(arr4)
}
