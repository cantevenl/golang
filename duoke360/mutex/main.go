package main

import (
	"fmt"
	"sync"
)

var i int = 100

var wg sync.WaitGroup

//互斥锁
var lock sync.Mutex

func add() {
	defer wg.Done()
	lock.Lock()
	i += 1
	fmt.Println("i++：", i)
	lock.Unlock()
}

func sub() {
	lock.Lock()
	defer wg.Done()
	i -= 1
	fmt.Println("i--：", i)
	lock.Unlock()
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
		wg.Add(1)
		go sub()
	}
	wg.Wait()

	fmt.Println("end i：", i)
}
