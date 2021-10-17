package main

import "fmt"

func main() {
	var arr [4][6]int
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	//遍历二维数组
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println()
	var arr2 [2][3]int //以这个为例来分析arr2在内存中的布局
	arr2[1][1] = 10
	fmt.Println(arr2)

	fmt.Printf("arr2[0]的地址%p\n", &arr2[0])
	fmt.Printf("arr2[1]的地址%p\n", &arr2[1])

	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][0])
	fmt.Printf("arr2[0][1]的地址%p\n", &arr2[0][1])

	var arr3 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var arr4 [2][3]int = [...][3]int{{2, 3, 4}, {3, 4, 5}}
	var arr5 = [3][3]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	var arr6 = [...][3]int{{1, 2, 3}, {2, 3, 4}, {3, 10, 5}}
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)
	fmt.Println(arr6)

	for i := 0; i < len(arr6); i++ {
		for j := 0; j < len(arr6[i]); j++ {
			fmt.Printf("%v\t", arr6[i][j])
		}
		fmt.Println()
	}

	for i, v := range arr6 {
		for i1, v1 := range v {
			fmt.Printf("arr6[%v][%v]=%v\t", i, i1, v1)
		}
		fmt.Println()
	}

}
