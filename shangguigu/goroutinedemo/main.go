package main

import (
	"fmt"
	"strconv"
	"time"
)

//开启一个goroutine，以函数为基本开启
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() //开启了一个协程
	//在主线程中也每隔一秒输入“hello,golang”
	for i := 1; i <= 5; i++ {
		fmt.Println("hello,golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
