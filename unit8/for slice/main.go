package main

import "fmt"

func main() {
	slice := make([]int, 4, 20)
	slice[0] = 66
	slice[1] = 88
	slice[2] = 99
	slice[3] = 100
	//方式1：for循环
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v\t", i, slice[i])
	}
	//方式2：for range
	fmt.Println("\n------------------------")
	for i, v := range slice {
		fmt.Printf("下标为%v,元素为%v\n", i, v)

	}

}
