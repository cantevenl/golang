package main

import "fmt"

//冒泡排序
func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr=", (*arr))
	for i := 0; i < len(*arr)-1; i++ {

		//完成第1轮排序(外层1次)
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				//交换
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]

			}
		}
		fmt.Printf("第%v次排序后arr=%v\n", i, arr)
	}

}

func main() {
	//定义数组
	arr := [5]int{27, 9, 88, 25, 5}
	//将数组传递给一个函数，完成排序
	BubbleSort(&arr)

	fmt.Println("main函数的arr=", arr) //有序的。因为传入的地址
}
