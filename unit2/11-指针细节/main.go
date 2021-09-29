package main

import "fmt"

func main() {
	var num int = 10
	fmt.Println(num)

	//var ptr *int = num  必须是地址值
	var ptr *int = &num
	*ptr = 20
	fmt.Println(num)

	//var ptr *float32 = &num	指针变量的地址必须匹配
	//*float32意味着这个指针指向的是float32类型的数据，但是&num对应的是int类型，因此会报错

}
