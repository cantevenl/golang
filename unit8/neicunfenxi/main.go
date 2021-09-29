package main

import "fmt"

func main() {
	var intarr [6]int = [6]int{3, 6, 9, 1, 4, 7}
	//切片构建在数组之上
	//定义一个切片名字为slice，[]动态变化的数组，长度不写，int是类型
	//[1:3]切片 表示 索引从1开始到3，不包含3
	slice := intarr[1:3]
	fmt.Println("intarr:", intarr)
	fmt.Println("slice:", slice)
	fmt.Println("slice元素个数：", len(slice))
	fmt.Println("slice的容量：", cap(slice))
	fmt.Printf("数组中下标为1位置的地址：%p\n", &intarr[1])
	fmt.Printf("切片中下标为0位置的地址：%p", &slice[0])

	//通过切片改变数组的元素
	slice[1] = 88
	fmt.Println("intarr:", intarr)
	fmt.Println("slice:", slice)

}
