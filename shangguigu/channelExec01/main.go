package main

import (
	"fmt"
	"time"
)

func main() {
	numChan := make(chan int, 1000)
	resChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)
	num := 0
	num2 := 0
	//统计时间
	start := time.Now().UnixMilli()
	go func(numChan chan int) {
		for i := 1; i <= 1000; i++ {
			//放入数据
			numChan <- i
		}
		close(numChan)
	}(numChan)

	for i := 1; i <= 8; i++ {
		go func(numChan chan int, resChan chan int, exitChan chan bool) {
			for {
				v, ok := <-numChan
				if !ok {
					break
				}
				num += v
				resChan <- num
			}
			exitChan <- true
		}(numChan, resChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		end := time.Now().UnixMilli()
		fmt.Println("使用协程耗时=", end-start)
		close(resChan)
	}()

	for {
		res, ok := <-resChan
		if !ok {
			break
		}
		num2++
		fmt.Printf("[%v]值为%v\n", num2, res)

	}
}
