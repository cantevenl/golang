package main

import "fmt"

func main() {
	//演示一下管道的使用
	//1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	//2.看看intChan是什么
	fmt.Printf("intChan的值=%v，intChan本身的地址=%p\n", intChan, &intChan)

	//3.向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	//注意点，当我们给管道写入数据时，不能超过其容量cap
	intChan <- 50
	//intChan <- 23	fatal error: all goroutines are asleep - deadlock!

	//4.看看管道的长度和cap(容量)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan)) //channel len=3 cap=3

	//5. 从管道中读取数据
	var num2 int
	num2 = <-intChan
	fmt.Printf("num2=%v\n", num2)
	//取了一个只会，len变少了
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan)) //channel len=2 cap=3

	//6.在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	num3 := <-intChan
	num4 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4)
	fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan)) //channel len=0 cap=3

	//再取一个就报错
	//但是如果继续存放一个就不会报错
	intChan <- 666
	num5 := <-intChan
	fmt.Println(num5)

}
