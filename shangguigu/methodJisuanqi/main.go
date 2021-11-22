package main

import "fmt"

/*
定义小小计算器结构体(Calcuator)
实现加减乘除
实现形式1：分四个方法完成：分别计算加减乘除
实现形式2：用一个方法搞定，需要接收两个数，还要一个运算符号
*/

//实现形式1
type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (calcuator *Calcuator) getSum() float64 {
	return calcuator.Num1 + calcuator.Num2
}

func (calcuator *Calcuator) getSub() float64 {
	return calcuator.Num1 - calcuator.Num2
}

//实现形式2
func (calcuator *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("运算符输入有误")

	}
	return res

}

func main() {
	var calcuator Calcuator
	calcuator.Num1 = 1.2
	calcuator.Num2 = 2.2
	fmt.Printf("sum=%v\n", fmt.Sprintf("%.2f", calcuator.getSum()))
	fmt.Printf("sub=%v\n", fmt.Sprintf("%.2f", calcuator.getSub()))

	res := calcuator.getRes('*')
	fmt.Println("res=", res)
}
