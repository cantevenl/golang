package main

import (
	"fmt"
	"sync/atomic"
)

func test_load_store() {
	var i int32=100
	atomic.LoadInt32(&i)	//read
	fmt.Printf("i: %v\n",i)

	atomic.StoreInt32(&i,200)	//write
	fmt.Printf("i: %v\n",i)
}

func test_cas(){
	//cas
	var i int32 =100
	b:=atomic.CompareAndSwapInt32(&i,100,200)
	fmt.Println("b：",b)
	fmt.Println("i：",i)
}

func main() {
	test_load_store()
	test_cas()
}
