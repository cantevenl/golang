package main

import "fmt"

func main() {
	//定义一个数组：
	var intarr [6]int = [6]int{3, 6, 9, 1, 4, 7}
	//切片构建在数组之上
	//定义一个切片，名字是slice，[]动态变化的数组，长度不写，int是类型
	//[1:3]切片 - 切出的一段片段 - 索引：从1开始，到3结束（不包含3）左闭右开
	//var slice []int = intarr[1:3]
	slice := intarr[1:3]
	fmt.Println(slice)
	//切片元素个数
	fmt.Println("slice元素个数是：", len(slice))
	//获取切片的容量:容量可以动态变化
	fmt.Println("slice的容量是：", cap(slice))
}
