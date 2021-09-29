package main

import "fmt"

func main() {
	//定义二维数组
	var arr [3][3]int = [3][3]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
	fmt.Println(arr)
	fmt.Println("-------------------------------------")

	//方式1：普通for循环
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Println()
	}

	fmt.Println("----------------------------------------")
	//方式2：for range
	for k, v := range arr {
		for k1, v1 := range v {
			fmt.Printf("arr[%v][%v]=%v", k, k1, v1)
		}
		fmt.Println()
	}
}
