package main

import (
	"fmt"
	"time"
)

func test() {
	//使用defer+recover来捕获和处理异常
	defer func() {
		err := recover() //recover()是一个内置函数，可以捕获到异常
		if err != nil {  //说明捕获到错误
			fmt.Println("err=", err)
			//这里可以将错误信息发给管理员
			fmt.Println("将错误发送给admin")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main() {
	test()
	for {
		fmt.Println("main()下面的代码...")
		time.Sleep(time.Second)
	}
}
