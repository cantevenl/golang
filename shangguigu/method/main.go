package main

import "fmt"

type Person struct {
	Name string
}

//给Person类型绑定一个方法
func (p Person) test() {
	p.Name = "猫"
	fmt.Println("test()", p.Name)
}

//给Person结构体添加speak方法，输出xxx是一个好人
func (p Person) speak() {
	fmt.Println(p.Name, "是一个好人")
}

//给Person结构体添加jisuan方法，计算1+...+1000的结果
func (p Person) jisuan() {
	res := 0
	for i := 0; i <= 1000; i++ {
		res += i
	}
	fmt.Printf("这个%v计算的结果是%v\n", p.Name, res)
}

//给Person结构体添加jisuan2方法，该方法可以接收一个数n，计算从1+...+n的结果
func (p Person) jisuan2(n int) {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	fmt.Printf("这个%v计算的结果是%v\n", p.Name, res)
}

//给Person结构体添加getSum方法，可以计算两个数的和，并返回结果
func (p Person) getSum(a, b int) int {
	return a + b
}

func main() {
	var p Person
	p.Name = "猪猪"
	p.test() //调用方法
	fmt.Println("main() p.Name=", p.Name)

	p.speak()
	p.jisuan()
	p.jisuan2(100)
	res := p.getSum(9, 27)
	fmt.Println("res=", res)
}
