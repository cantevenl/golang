package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)
var wg sync.WaitGroup
var i int32 = 100

func add() {
	defer wg.Done()
	atomic.AddInt32(&i, 1)
}

func sub() {
	defer wg.Done()
	atomic.AddInt32(&i, -1)
}

func main() {

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait()
	fmt.Printf("i: %v\n", i)
}