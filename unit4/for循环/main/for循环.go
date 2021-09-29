package main

import "fmt"

func main() {
	//求和：1+2+3+...+100
	sum := 0
	//for循环语法：
	//for 初始表达式;布尔表达式（条件判断）;迭代因子 {
	//	循环体
	//}
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	//for的初始表达式 不能用var定义变量的形式，要用:=
	//for循环实际就是让程序员写代码的效率变高，但是底层该怎么执行还是怎么执行，底层效率没有提高
}
