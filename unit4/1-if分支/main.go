package main

import "fmt"

func main() {
	//实现功能：如果口罩库存小于30个，提示：库存不足
	var count int
	fmt.Println("请输入口罩数：")
	//fmt.Scanln(&count)
	//单分支
	//if count < 30 {
	//	fmt.Println("对不起口罩存储不足")
	//}

	//if后面表达式，返回结果一定是true或者false
	//如果返回结果为true才会执行{}里面的代码，如果返回结果是false，那么{}中的代码就不会执行
	//if后面一定要有空格，和条件表达式分割开来

	if fmt.Scanln(&count); count < 20 {
		fmt.Println("对不起没了")
	}
}
