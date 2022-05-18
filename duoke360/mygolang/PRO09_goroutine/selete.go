package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int, 0)
var chanStr = make(chan string)

func main() {
	go func() {
		//defer close(chanInt)
		//defer close(chanStr)
		chanInt <- 100
		chanStr <- "hello"
	}()

	for {
		select {
		case r := <-chanInt:
			fmt.Println(r)
		case r := <-chanStr:
			fmt.Println(r)
		default:
			fmt.Println("default.....")

		}
		time.Sleep(1 * time.Second)
	}
}
