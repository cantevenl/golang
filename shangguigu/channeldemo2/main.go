package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	//定义一个存放任意数据类型的管道 3个数据
	//var allChan chan interface{}
	allChan := make(chan interface{}, 3)
	allChan <- 88
	allChan <- "牛"
	cat := Cat{"喵喵", 1}
	allChan <- cat

	//我们希望获得管道中第三个元素，则先将前两个推出
	<-allChan
	<-allChan

	newCat := <-allChan //从管道里取出的Cat是什么

	fmt.Printf("newCat的类型=%T,newCat=%v\n", newCat, newCat)
	//下面写法是错误的，编译不通过
	//fmt.Printf("newCat.Name=%v",newCat.Name)
	//使用类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v\n", a.Name)
}
