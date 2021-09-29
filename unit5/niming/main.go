package main

import "fmt"

var Func01 = func(num1 int, num2 int) int {
	return num1 * num2
}

func main() {
	//定义匿名函数：
	result := func(num1 int, num2 int) int {
		return num1 + num2
	}(27, 25)
	fmt.Println(result)

	//将匿名函数赋给一个变量，这个变量实际就是函数类型的变量
	//sub等价与匿名函数
	sub := func(num1 int, num2 int) int {
		return num1 - num2
	}

	//直接调用sub
	result02 := sub(30, 70)
	fmt.Println(result02)

	result03 := sub(70, 99)
	fmt.Println(result03)

	result04 := Func01(7, 9)
	fmt.Println(result04)
}
