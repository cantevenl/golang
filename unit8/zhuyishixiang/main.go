package main

import "fmt"

func main() {
	//定义数组
	var intarr [6]int = [6]int{1, 4, 7, 2, 5, 8}
	//定义一个切片：
	var slice []int = intarr[1:4] // 4 7 2
	//对切片再次切片
	slice2 := slice[1:2] // 7
	fmt.Println(slice2)
	slice2[0] = 66
	fmt.Println(intarr)
	fmt.Println(slice)

	//底层原理：
	//1.底层追加元素的时候对数组进行扩容，老数组扩容为新数组
	//2.创建一个新数组，将老数组中的4,7,2复制到新数组中，在新数组中追加88,55
	//3.slice3的底层数组的指向，指向的是新数组
	//4.往往我们在使用追加的时候其实想要把效果给slice追加
	slice3 := append(slice, 88, 270)
	fmt.Println(slice3)
	fmt.Println(slice)
	slice = append(slice, 88, 270)
	fmt.Println(slice)
	//5.底层的新数组 不能直接维护，需要通过切片间接维护操作

	slice4 := []int{99, 44}
	slice = append(slice, slice4...)
	fmt.Println(slice)
}
