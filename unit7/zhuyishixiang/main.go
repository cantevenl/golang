package main

import "fmt"

func main() {
	//定义一个数组：
	var arr1 = [3]int{3, 6, 9}
	fmt.Printf("数组的类型是：%T\n", arr1)

	var arr2 = [6]int{1, 2, 3, 4, 6, 9}
	fmt.Printf("数组的类型是：%T\n", arr2)

	var arr3 = [3]int{3, 6, 9}
	test1(arr3)
	fmt.Println(arr3) //3, 6, 9
	test2(&arr3)
	fmt.Println("改变后的arr3：", arr3)
}

func test1(arr [3]int) {
	arr[0] = 7
}

func test2(arr *[3]int) {
	(*arr)[0] = 7
}
