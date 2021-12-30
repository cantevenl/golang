package main

import (
	"fmt"
	"math/rand"
	"time"
)

//随机生成10个整数保存到数组，并倒序打印以及求平均值、最大值和最大值下标、并且查找里面是否有55

//生成随机数组的方法
func Random(arr *[10]int) [10]int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] = rand.Intn(100) //0<=n<=100
	}
	//fmt.Printf("反转前：%v\n",arr)
	//for i:=0;i<len(arr)/2;i++ {
	//	arr[i],arr[len(arr)-1-i] = arr[len(arr)-1-i],arr[i]
	//}
	//fmt.Printf("反转后：%v",arr)
	return *arr
}

//正序打印

func BubbleSort1(arr *[10]int) {
	fmt.Println("排序前arr=", (*arr))
	for i := 0; i < len(*arr)-1; i++ {
		flag := true //假设排序没有进行比较时为true
		for j := 0; j < len(*arr)-1-i; j++ {
			//fmt.Printf("(*arr)[j]=%v,(*arr)[j+1]=%v\n",(*arr)[j],(*arr)[j+1])
			if (*arr)[j] > (*arr)[j+1] {
				flag = false //进行比较则变成false
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
		if flag {
			//如果没有比较就break跳出循环
			break
		}
		fmt.Printf("第%v次排序后arr=%v\n", i, *arr)
	}
}

//倒序打印并求平均值
func BubbleSort2(arr *[10]int) [10]int {
	fmt.Println("排序前arr=", (*arr))
	for i := 0; i < len(*arr)-1; i++ {
		flag := true //假设排序没有进行比较时为true
		for j := len(*arr) - 1; j > 0; j-- {
			//fmt.Printf("(*arr)[j]=%v,(*arr)[j+1]=%v\n",(*arr)[j],(*arr)[j+1])
			if (*arr)[j] > (*arr)[j-1] {
				flag = false //进行比较则变成false
				(*arr)[j], (*arr)[j-1] = (*arr)[j-1], (*arr)[j]
			}
		}
		if flag {
			//如果没有比较就break跳出循环
			break
		}
		fmt.Printf("第%v次排序后arr=%v\n", i, *arr)

	}
	return *arr
}

//求平均值
func Aug(arr *[10]int) {
	sum := 0
	for _, v := range *arr {
		sum += v
	}
	fmt.Printf("数组内元素的平均值是：%v\n", sum/len(*arr))
}

//查找
func Find(arr *[10]int, leftIndex, rightIndex, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("没有找到")
		return
	}

	midIndex := (leftIndex + rightIndex) / 2
	if (*arr)[midIndex] > findVal {
		//说明我们要查的数，应该在 leftIndex 到 midIndex-1 之间
		Find(arr, leftIndex, midIndex-1, findVal)
	} else if (*arr)[midIndex] < findVal {
		//说明我们要查的数，应该在 midIndex+1 到 rightIndex 之间
		Find(arr, midIndex+1, rightIndex, findVal)
	} else {
		//找到了
		fmt.Printf("找到了%v下标为%v", findVal, midIndex)
	}
}

func main() {
	var arr [10]int
	Random(&arr)
	BubbleSort2(&arr)
	Aug(&arr)
	Find(&arr, 0, len(arr)-1, 55)

}
