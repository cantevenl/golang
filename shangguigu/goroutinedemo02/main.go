package main

import (
	"fmt"
	"sync"
	"time"
)

//现在要求计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中。
//最后显示出来，要求使用goroutine完成

//思路
//1.编写一个函数，来计算各个数的阶乘，并放入到map中
//2.我们启动多个协程，统一的将结果放入到map中
//3.map应该是一个全局的

var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock是一个全局的互斥锁
	//sync是包：表示同步的意思
	//Mutex是互斥
	lock sync.Mutex
)

//test函数计算n!，将这个结果放入map
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	//这里将res放入myMap
	//加锁
	lock.Lock()
	myMap[n] = res
	//解锁
	lock.Unlock()
}

func main() {
	//我们这里开启多个协程完成这个任务[200个]
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 3)

	//这里我们输出结果,遍历这个结果
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("myMap[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
