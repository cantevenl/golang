//package main
//
//import "fmt"
//
////函数功能求和
////函数名字为getSum 参数为空
////getSum函数返回值为一个函数，这个函数的参数是一个int类型的参数，返回值也是int类型
//func getSum() func(int) int {
//	var sum int = 0
//	return func(num int) int {
//		sum = sum + num
//		return sum
//	}
//}
//
//func main() {
//	f := getSum()
//	fmt.Println(f(1))
//	fmt.Println(f(2))
//	fmt.Println(f(3))
//}

package main

import "fmt"

//求和
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum += num
		return sum
	}
}

func main() {
	f := getSum()
	fmt.Println(f(11))
	fmt.Println(f(22))

}
