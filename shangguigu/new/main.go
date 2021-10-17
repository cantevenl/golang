package main

import "fmt"

func main() {
	num1 := 100
	fmt.Printf("num1的类型是%T,num1的值是%+v,num1的地址是%v\n", num1, num1, &num1)

	num2 := new(int) //*int
	//num2的类型是*int
	//num2的值是一个地址
	fmt.Printf("num2的类型%T,num2的值是%v,num2的地址%v,num2这个指针指向的值是%v\n", num2, num2, &num2, *num2)
	*num2 = 100
	fmt.Println("num2这个指针指向的值是:", *num2)
}
