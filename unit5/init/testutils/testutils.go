package testutils

import "fmt"

var Age int
var Gender string
var Name string

//定义一个init函数，对变量进行初始化赋值：
func init() {
	fmt.Println("testutils中的init函数被执行了")
	Age = 19
	Gender = "女"
	Name = "喵喵"
}
