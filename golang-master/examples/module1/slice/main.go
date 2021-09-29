package main

import (
	"fmt"
)

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	mySlice := myArray[1:3]
	mySlice1 := []int{1, 3, 4}
	fmt.Println(mySlice1)
	mySlice1 = append(mySlice1, 27)
	mySlice1 = append(mySlice1, 25)
	mySlice1 = append(mySlice1, 88)
	fmt.Println("mySlice1：", mySlice1)
	fmt.Printf("mySlice %v\n", mySlice)
	fullSlice := myArray[:]
	//删除第3个下标的元素
	remove3rdItem := deleteItem(fullSlice, 3)
	fmt.Printf("remove3rdItem %v\n", remove3rdItem)
}

//删除切片
func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
