package main

import (
	"fmt"
	"unsafe"
)

//全局变量，定义在函数外的变量
var m1 = 100
var m2 = 9.124

//设计者认为上面的全局变量写法太麻烦，可以一次性声明：
var (
	m3 = "牛"
	m4 = 3306
)

func main() {
	//第一种：变量的使用方式，指定变量的类型，并且赋值
	var num = 18
	fmt.Println(num)

	//第二种：指定变量的类型，但是不赋值，使用默认值
	var num2 int
	fmt.Println(num2)

	//第三种：如果没有写变量的类型，那么根据=后面的值进行判定变量的类型（自动类型判断）
	var num3 = "niu"
	fmt.Println(num3)

	//第四种：省略var，是:=
	gender := "男"
	fmt.Println(gender)

	fmt.Println("---------------------------------------------------------------")
	//多变量声明
	var n1, n2, n3 = 1, 2, 3
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)

	var n4, n5, n6 = 4, "牛", "niu"
	fmt.Println(n4)
	fmt.Println(n5)
	fmt.Println(n6)

	n7, n8, n9 := 6.8, "猪", "鸡"
	fmt.Println(n7)
	fmt.Println(n8)
	fmt.Println(n9)

	fmt.Println(m1, m2, m3, m4)

	zhu := 28
	fmt.Printf("zhu的类型是: %T", zhu)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(zhu))

	//表示学生年龄
	var age uint8 = 18
	var age1 byte = 18
	fmt.Println(age, age1)

}
