package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	//如何将a赋给一个Point变量
	var b Point
	//b = a不可以
	//b = a.(Point)
	b = a.(Point)
	fmt.Println(b)

	////类型断言的其他案例
	//var x interface{}
	//var b2 float32 = 1.1
	//x = b2 //空接口，可以接收任意类型
	//// x==>float32 [使用类型断言]
	//y := x.(float32)
	//fmt.Printf("y的类型是 %T 值是=%v", y, y)

	//待检测的类型断言
	var x interface{}
	var b2 float32 = 1.1
	x = b2 //空接口，可以接收任意类型
	// x==>float32 [使用类型断言]
	//y,flag:= x.(float64)
	//if flag==true {
	//	fmt.Println("转换成功")
	//	fmt.Printf("y的类型是 %T 值是=%v", y, y)
	//}else {
	//	fmt.Println("转换失败")
	//}
	//
	//fmt.Println("继续执行")

	//y,ok:= x.(float64)
	//if ok {	//ok本身就是true
	//	fmt.Println("转换成功")
	//	fmt.Printf("y的类型是 %T 值是=%v", y, y)
	//}else {
	//	fmt.Println("转换失败")
	//}
	//
	//fmt.Println("继续执行")

	//简洁写法
	if y, ok := x.(float64); ok { //ok本身就是true
		fmt.Println("转换成功")
		fmt.Printf("y的类型是 %T 值是=%v", y, y)
	} else {
		fmt.Println("转换失败")
	}

	fmt.Println("继续执行")
}
