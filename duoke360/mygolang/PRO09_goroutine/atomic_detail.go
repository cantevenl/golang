package main

import (
	"fmt"
	"sync/atomic"
)

//修改
func test_add_sub() {
	var i int32 = 100
	atomic.AddInt32(&i, 1)
	fmt.Printf("i:%v\n", i)
	atomic.AddInt32(&i, -1)
	fmt.Printf("i:%v\n", i)

	var j int64 = 200
	atomic.AddInt64(&j, 10)
	fmt.Printf("j：%v\n", j)
	atomic.AddInt64(&j, -10)
	fmt.Printf("j：%v\n", j)
}

//读写
func test_load_store() {
	var i int32 = 100
	atomic.LoadInt32(&i)
	fmt.Printf("i：%v\n", i)

	atomic.StoreInt32(&i, 200)
	fmt.Printf("i：%v\n", i)
}

func test_cas() {
	//cas
	var i int32 = 100
	b := atomic.CompareAndSwapInt32(&i, 100, 200)
	fmt.Printf("b:%v\n", b)
	fmt.Printf("i:%v\n", i)
}

func main() {


}
