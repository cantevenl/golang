package main

import "fmt"

func main() {
	var age int = 18
	// &符号+变量 就可以获取到这个变量的内存地址
	fmt.Println(&age) //0xc00000a0b8

	//定义一个指针变量
	//var 代表要声明一个变量
	//ptr指针变量的名字
	//ptr对应的类型是：*int 是一个指针类型（可以理解为 指向int类型的指针）
	//&age就是一个地址：是ptr变量的具体的值
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr本身的内存地址为：", &ptr)

	//想获取ptr这个指针或者这个内存地址指向的那个数据
	fmt.Println("ptr指向的数值为：", *ptr)
	fmt.Printf("ptr指向的数值为：%v", *ptr)
}
