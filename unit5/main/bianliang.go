package main

import "fmt"

func test2(num int) {
	fmt.Println(num)
}

//定义一个函数，把另一个函数作为形参：
func test3(num1 int, num2 float32, testFunc func(int)) {
	fmt.Println("调用了函数test2")
}
func main() {
	//函数也是一种数据类型，可以赋值给一个变量
	a := test2                                         //变量就是一个函数类型的变量
	fmt.Printf("a对应的类型是:%T，test函数的类型是:%T\n", a, test2) //a对应的类型是:func(int)，test函数的类型是:func(int)

	//通过该变量可以对函数调用
	a(99) //等价于 test2(99)

	//调用test3函数：
	test3(27, 3.14, test2)
	test3(27, 3.14, a)

	//自定义数据类型：(相当于起别名)：给int类型起了别名叫myInt类型
	type myInt int

	var niu myInt = 30
	fmt.Println(niu)

	var niu1 int = 60

	niu1 = int(niu)
	fmt.Println(niu1)
}
