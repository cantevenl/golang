package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1 int = 19
	var n2 float32 = 4.78
	var n3 bool = false
	var n4 byte = 'a'

	var s1 string = fmt.Sprintf("%d", n1)
	fmt.Printf("s1的类型是：%T , s1 = %q ", s1, s1)

	var s2 string = fmt.Sprintf("%F", n2)
	fmt.Printf("s2的类型是：%T , s2 = %q ", s2, s2)

	var s3 string = fmt.Sprintf("%t", n3)
	fmt.Printf("s3的类型是：%T , s3 = %q ", s3, s3)

	var s4 string = fmt.Sprintf("%c", n4)
	fmt.Printf("s4的类型是：%T, s4 = %q ", s4, s4)

	fmt.Println("\n---------------------------------------------------")

	var m1 int = 18
	//参数：第一个参数必须转为int64类型，第二个参数指定字面值的进制形式为十进制
	var a1 string = strconv.FormatInt(int64(m1), 10)
	fmt.Printf("a1的类型是：%T, a1 = %q", a1, a1)

	var m2 float64 = 3.1415926
	var a2 string = strconv.FormatFloat(m2, 'f', 9, 64)
	fmt.Printf("a2的类型是：%T, a2 = %q", a2, a2)

	var a3 string = fmt.Sprintf("%F", m2)
	fmt.Printf("a3的类型是：%T, a3 = %q", a3, a3)

	var m3 bool = true
	var a4 string = strconv.FormatBool(m3)
	fmt.Printf("s3的类型是：%T , a4 = %q", a4, a4)
}
