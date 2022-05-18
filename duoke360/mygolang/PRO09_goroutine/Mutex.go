package main

import (
	"fmt"
	"sync"
	"time"
)

var m int = 100

var lock sync.Mutex

var wt sync.WaitGroup

func add() {
	defer wt.Done()
	lock.Lock()
	m += 1
	time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func sub() {
	defer wt.Done()
	lock.Lock()
	time.Sleep(time.Millisecond * 2)
	m -= 1
	lock.Unlock()
}

func main() {

	for i := 0; i < 100; i++ {
		go add()
		wt.Add(1)
		go sub()
		wt.Add(1)
	}

	wt.Wait()
	fmt.Printf("m: %v\n", m)
}
