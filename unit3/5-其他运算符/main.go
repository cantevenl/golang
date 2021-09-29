package main

import "fmt"

func main() {
	var age int = 10
	fmt.Println("age对应的内存地址为：", &age)

	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr这个指针指向的具体数值为:", *ptr)
}
