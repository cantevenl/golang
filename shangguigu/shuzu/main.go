package main

import (
	"fmt"
	"math/rand"
	"time"
)

//要求：随机生成五个数，并将其反转打印
//1.随机生成五个数，rand.Intn()函数
//2.当我们得到随机数后，就放到一个数组 int数组
//3.反转打印,交换的次数是 len/2, 倒数第一个和第一个元素交换，倒数第二个和第二个交换
func main() {
	var intArr [5]int
	//为了每次生成的随机数不一样，需要给一个seed值
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100) // 0<=n<100
	}

	fmt.Println("交换前=", intArr)
	temp := 0
	for i := 0; i < len(intArr)/2; i++ {
		temp = intArr[len(intArr)-1-i]
		intArr[len(intArr)-1-i] = intArr[i]
		intArr[i] = temp
	}
	fmt.Println("交换后=", intArr)

}
