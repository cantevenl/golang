package main

import "fmt"

//加法函数
//func 函数名 (形参列表) (返回值类型列表) {
// 执行语句...
// return + 返回值列表
// }
func cal(num1 int, num2 int) int { //如果返回值类型就一个的话，那么()是可以省略不写的
	sum := 0
	sum += num1 + num2
	return sum
}

func main() {
	sum := cal(27, 25)
	fmt.Println(sum)

	var num3 int = 30
	var num4 int = 50
	sum1 := cal(num3, num4)
	fmt.Println(sum1)
}
