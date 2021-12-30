package main

//go 中，每个可运行程序都有一个主package。它用来告诉Go这个是一个可运行程序，最终可以编译产生可以直接运行的程序

import "fmt"

//func 开头表示这是一个 function，
//翻译为“方法”或“函数”。main
//package 里边的 main 函数，是在
//程序运行时的起始点

func main() { //两个花括号中间的部分，为函数主体。主体里边的内容按顺序从上到下执行。
	fmt.Println("你好,golang!")

	var hello string = "hello!golang!"
	fmt.Println(hello)

	var age int = 35
	var tall float64 = 1.70
	fmt.Println(age, tall)

}
