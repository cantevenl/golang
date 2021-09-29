package main

import "fmt"

func main() {
	//声明数组：
	var arr [3]int16
	//获取数组长度
	print(len(arr))
	//打印数组：
	fmt.Println(arr)
	//证明arr中存储的值是地址值：
	fmt.Printf("arr的地址为：%p\n", &arr)
	//第一个空间的地址：
	fmt.Printf("arr的地址为：%p\n", &arr[0])
	fmt.Printf("arr的地址为：%p\n", &arr[1])
	fmt.Printf("arr的地址为：%p\n", &arr[2])

	//赋值：
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

}
