package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func show() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		if i > 5 {
			runtime.Goexit()
		}
		fmt.Printf("i: %v\n", i)
	}
}

func main() {
	wg.Add(1)
	go show()
	wg.Wait()
}