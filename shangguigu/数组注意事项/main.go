package main

import "fmt"

func test(arr [3]int) {
	fmt.Printf("test函数内arr的地址：%p\n", &arr[0])
	arr[0] = 88
}

func test1(arr *[3]int) {
	fmt.Printf("test1函数内arr指针的地址：%p\n", &arr[0])
	(*arr)[0] = 88
}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Printf("arr的地址：%p\n", &arr)
	test(arr)
	fmt.Println(arr)
	test1(&arr)
	fmt.Println(arr)
	fmt.Printf("arr的第一个元素地址：%p\n", &arr[0])
}
