package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//方式2
	p := Person{"猪", 2}
	fmt.Println(p)

	//方式3
	// var p1 *Person = new(Person)
	p1 := new(Person)
	//因为p1是一个指针，因此标准的给字段赋值的方式是：
	(*p1).Name = "猫"
	(*p1).Age = 1
	fmt.Printf("这是一只%v岁的%v\n", (*p1).Age, (*p1).Name)

	//(*p1).Name="猫"也可以这样写 p1.Name = "猫"
	//原因是：go的设计者为了程序员使用方便，底层会对p1.Name = "猫" 进行处理
	//会给p1加上取至运算
	p1.Name = "猫猫"
	p1.Age = 2
	fmt.Printf("这是一只%v岁的%v\n", (*p1).Age, (*p1).Name)

	//方式4
	var p2 *Person = &Person{"牛", 5}
	fmt.Printf("这是一只%v岁的%v", (*p2).Age, (*p2).Name)

	//因为p2是一个指针，因此标准访问字段的方法
	//(*p2).Name = "牛"
	//go的设计者为了程序员使用方便，也可以写p2.Name = "牛"
}
