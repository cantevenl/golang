package main

import "fmt"

func main() {
	//go语言也有指针
	//结构体成员调用时：	c语言：	ptr->name	go语言：	ptr.name

	//go语言在使用指针时，会使用内部的垃圾回收机制（gc ： garbage collector），开发人员不需要手动释放
	//c语言不允许返回栈上的指针，go语言可以返回栈上的指针，程序会在编译的时候就确定了变量的分配空间
	//编译对的时候，如果发现有必要的话，就将变量分配到栈上

	name := "孙子"
	ptr := &name
	fmt.Println("name:", *ptr)
	fmt.Println("ptr", ptr)

}
