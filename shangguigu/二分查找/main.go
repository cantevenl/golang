package main

import "fmt"

func BinaryFind(arr *[]int, leftIndex, rightIndex, findVal int) {

	//判断leftIndex是否大于rightIndex
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	//先找到中间的下标
	midIndex := (leftIndex + rightIndex) / 2
	if (*arr)[midIndex] > findVal {
		//说明我们要查的数，应该在 leftIndex 到 midIndex-1 之间
		BinaryFind(arr, leftIndex, midIndex-1, findVal)
	} else if (*arr)[midIndex] < findVal {
		//说明我们要查的数，应该在 midIndex+1 到 rightIndex 之间
		BinaryFind(arr, midIndex+1, rightIndex, findVal)
	} else {
		//找到了
		fmt.Printf("找到了下标为%v \n", midIndex)
	}
}

func main() {
	arr := []int{3, 8, 9, 25, 27, 66, 88, 230, 520, 1314, 8888}
	fmt.Println("输入你要查找的数：")
	var num int
	fmt.Scanln(&num)
	BinaryFind(&arr, 0, len(arr)-1, num)
}
