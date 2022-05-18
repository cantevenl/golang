package main

import (
	"fmt"
	"math/rand"
	"time"
)

var values = make(chan int)

var c = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send:%v\n", value)
	time.Sleep(time.Second * 1)
	values <- value
}

func main() {
	defer close(values)
	go send()
	fmt.Println("wait...")
	value := <-values
	fmt.Printf("receive: %v\n", value)
	fmt.Println("end...")

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Printf("v:%v\n", v)
	}

	for i:=0;i<11;i++ {
		r := <-c
		fmt.Printf("r:%v\n",r)
	}

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c <- i
	//	}
	//	close(c)
	//}()
	//
	//for {
	//	if data, ok := <-c; ok {
	//		fmt.Printf("data: %v\n", data)
	//	} else {
	//		break
	//	}
	//}

}
