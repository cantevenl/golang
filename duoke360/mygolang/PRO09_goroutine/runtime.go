package main

import (
	"fmt"
	"runtime"
	"sync"
)

func show(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Println(msg,i)
	}
}

func a() {
	defer wp.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	defer wp.Done()
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

var wp sync.WaitGroup

func main() {
	go show("golang")
	for i := 0; i < 2; i++ {
		runtime.Gosched()	//让其他子协程来执行
		fmt.Println("k8s")
	}

	fmt.Println("end.........")

	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(24) // 修改为1查看效果
	wp.Add(1)
	go a()
	wp.Add(1)
	go b()
	wp.Wait()
	//time.Sleep(time.Second)
}

