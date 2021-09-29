package main

import "fmt"

func main() {
	var n1 int = 10
	fmt.Println(n1)

	var n2 int = (10 + 20*93/9) % 3 //等号右侧的值运算清楚之后，再赋值给=的左侧
	fmt.Println(n2)

	var n3 int = 10
	//n3 = n3 +20;
	n3 += 20
	fmt.Println(n3)

	var n4 float32 = 3.1415926
	n4 *= 3.14
	fmt.Println(n4)

	//交换两个数的值，并输出结果
	var a int = 8
	var b int = 92
	fmt.Printf("a=%v,b=%v\n", a, b)

	//交换
	//引入一个中间变量
	var t int
	t = a
	a = b
	b = t
	fmt.Printf("a=%v,b=%v", a, b)

}
