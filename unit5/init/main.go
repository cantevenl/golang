package main

import (
	"fmt"
	"unit5/init/testutils"
)

var num int = test()

func test() int {
	fmt.Println("test函数被执行")
	return 10
}

func init() {
	fmt.Println("main.go中的init函数被执行了")
}

func main() {
	fmt.Println("main函数被执行了")
	fmt.Printf("Age=%v,Sex=%v,Name=%v", testutils.Age, testutils.Gender, testutils.Name)

}
