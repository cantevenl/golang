package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}
	//关闭intChan
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	//使用for循环
	var flag bool
	for {
		//time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok { //intChan管道取不到数了
			break
		}
		flag = true // 假定是素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			//将这个数放入primeChan
			primeChan <- num
		}
	}
	//fmt.Println("有一个primeChan协程因为取不到数据退出")
	//这里还不能关闭primeChan这个管道
	//向exitChan写入true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 200000)
	primeChan := make(chan int, 1000000) //放入结果
	//标识退出的管道
	exitChan := make(chan bool, 8) //8个

	//统计时间
	start := time.Now().UnixMilli()

	//开启一个协程，向intChan放入1-200000个数
	go putNum(intChan)

	//开启八个协程，从intChan取出数据，并判断是否是素数，如果是素数就放入primeChan
	for i := 0; i < 8; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}

		end := time.Now().UnixMilli()
		fmt.Println("使用协程耗时=", end-start)
		//当我们从exitChan取出了8个结果，就可以放心的关闭primeChan
		close(primeChan)
	}()

	//遍历我们的primeChan，把结果取出
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//fmt.Printf("素数=%d\n", res)
	}
	fmt.Println("main线程退出")
}
