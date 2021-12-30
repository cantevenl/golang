package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//1.创建一个byte类型的26个元素的数组，分别放置'A'-'Z'
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}

	for _, v := range myChars {
		fmt.Printf("%c ", v)
	}

	//2.请求出一个数组的最大值，并得到其下标
	var intArr = [...]int{1, -1, 9, 88, 10}
	maxVal := intArr[0]
	maxIndex := 0

	for i := 0; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxIndex = i
		}
	}
	fmt.Printf("最大数是：%v，下标为：%v\n", maxVal, maxIndex)

	//3.求出数组的和和平均值
	var avg [7]int = [...]int{2, 4, 6, 8, 10, 11, 12}
	a := 0
	for _, v := range avg {
		a += v
	}
	//让值保留到小数位
	fmt.Println("平均值为：", float64(a)/float64(len(avg)))

	//4.数组反转，随机生成五个数，并将其反转打印
	//随机生产数，rand.Intn() 函数
	//得到随机数后放入数组
	//反转打印，交换的次数为 len/2 倒数第一个和第一个元素交换

	var intArr2 [5]int
	//为了每次生成的随机数不一样，我们需要给一个seed值
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(intArr2); i++ {
		intArr2[i] = rand.Intn(100) //0<=n<=100
	}
	fmt.Println("反转前：", intArr2)
	//反转打印，交换的次数为 len/2 倒数第一个和第一个元素交换
	for i := 0; i <= len(intArr2)/2; i++ {
		intArr2[i], intArr2[len(intArr2)-1-i] = intArr2[len(intArr2)-1-i], intArr2[i]
	}
	fmt.Println("反转后：", intArr2)
}

//反转前： [59 5 42 42 50]
//反转后： [50 42 42 5 59]
