package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var i int32 = 100

func add1() {
	atomic.AddInt32(&i, 1)
}

func sub1() {
	atomic.AddInt32(&i, -1)
}

func main() {
	for i := 0; i < 100; i++ {
		go add1()
		go sub1()
	}

	time.Sleep(time.Second * 3)
	fmt.Printf("i: %v\n", i)
}
