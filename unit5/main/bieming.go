package main

import "fmt"

type myFunc func(int)

func bieming(num1 int, num2 float32, myFunc myFunc) {
	fmt.Println("调用了函数")
}

func bie(num int) {
	fmt.Println("请输入数字：")
	fmt.Scanln(num)
	fmt.Println("你输入的数字是：", num)
}

//定义一个函数，求两个数的和 和 差
func re(num1 int, num2 int) (sum int, sub int) {
	sum = num1 + num2
	sub = num1 - num2
	return

}

func main() {
	a := bie
	bieming(1, 1.1, a)
	zhu, zhu1 := re(99, 88)
	fmt.Printf("结果是：%v,%v", zhu, zhu1)
}
