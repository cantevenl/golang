package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(24) // 修改为1查看效果
	go a()
	go b()
	time.Sleep(time.Second)
}