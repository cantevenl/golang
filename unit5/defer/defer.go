package main

import "fmt"

func main() {
	fmt.Println(add(30, 60))
}

func add(num1 int, num2 int) int {
	//在golang中，程序遇到defer关键字，不会立即执行defer后的语句，而是要将defer后的语句压入一个栈中，然后继续执行函数后面的语句
	defer fmt.Println("num1=", num1)
	defer fmt.Println("num2=", num2)
	//栈的特点是：先进后出
	//在这个函数执行完毕以后，从栈中取出语句开始执行。按照先进后出的规则执行这个语句
	var sum = num1 + num2
	fmt.Println("sum=", sum)
	return sum
}
