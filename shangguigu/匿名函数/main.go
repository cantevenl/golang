package main

import "fmt"

//全局匿名函数
var (
	//fun1就是一个全局匿名函数/
	Fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

func main() {
	res1 := func(n1, n2 int) int {
		return n1 + n2
	}(27, 9)

	fmt.Println(res1)

	//将匿名函数func (n1,n2 int)int 赋给了a变量
	//则a的数据类型就是函数类型，此时，我们可以通过a完成调用
	a := func(n1, n2 int) int {
		return n1 - n2
	}

	fmt.Printf("a的类型是%T\n", a)

	res2 := a(27, 9)
	fmt.Println(res2)

	res3 := a(90, 30)
	fmt.Println(res3)

	//全局匿名函数使用
	res4 := Fun1(27, 9)
	fmt.Println(res4)
}
