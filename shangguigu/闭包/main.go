package main

import (
	"fmt"
	"strings"
)

//累加器
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += string(x)
		fmt.Println("str=", str)
		return n

	}
}

func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果 name 没有指定的后缀，则加上，否则返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16

	//测试makeSuffix使用
	//返回一个闭包
	f1 := makeSuffix(".jpg")
	fmt.Println("文件名处理后=", f1("winter"))
	fmt.Println("文件名处理后=", f1("niu.tar"))
	fmt.Println("文件名处理后=", f1("zhu.jpg"))

}
