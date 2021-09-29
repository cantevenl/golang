package main

import "fmt"

func main() {
	//定义一个切片
	var a []int = []int{1, 4, 7, 3, 6, 9}
	//再顶一个切片
	var b []int = make([]int, 10)
	//拷贝：把a里元素给b
	copy(b, a) //将a中对应数组中元素内容复制到b中对应的数组中
	fmt.Println(b)

}
