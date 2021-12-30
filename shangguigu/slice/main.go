package main

import "fmt"

func main() {

	//var arr [3]int = [...]int{1, 2, 3}
	//fmt.Printf("arr地址：%p,arr[0]地址:%p", &arr, &arr[0])

	//var slice []int =[]int{1,2,3,4,5}
	//var slice []int = make([]int,5,10)
	//fmt.Println(slice)
	//fmt.Printf("扩容前slice地址：%p\n",&slice)
	//fmt.Println(slice)
	//fmt.Printf("扩容前slice[0]地址：%p\n",&slice[0])
	//
	//slice = append(slice,6,7,8)
	//fmt.Println(slice)
	//fmt.Printf("扩容1slice[0]地址：%p\n",&slice[0])
	//
	//slice = append(slice,slice...)
	//fmt.Println(slice)
	//fmt.Printf("扩容2slice[0]地址：%p\n",&slice[0])

	//切片make方法
	var slice []float64 = make([]float64, 5, 10)
	slice[1] = 10
	slice[3] = 20
	fmt.Println(slice)

	//定义一个切片，直接指定具体数组，使用原理类似make
	var strSlice []string = []string{"zhu", "niu", "yang"}
	fmt.Println("strSlice=", strSlice)
	fmt.Println("strSlice size=", len(strSlice))
	fmt.Println("strSlice cap=", cap(strSlice))

	var slice2 []int = []int{1, 2, 3}
	fmt.Println(slice2, &slice2[0])
	//通过append直接给slice2追加具体的元素
	slice2 = append(slice2, 4, 5, 6)
	fmt.Println(slice2, &slice2[0])
	//通过append将切片slice2追加给slice2
	slice2 = append(slice2, slice2...)
	fmt.Println(slice2, &slice2[0])

	var a []int = []int{1, 2, 3}
	var b = make([]int, 10)
	copy(b, a)
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	str := "hello@kubernetes"
	//使用切片获取到kubernetes
	slice3 := str[6:]
	fmt.Println(slice3)

	//如果需要修改字符串，可以先将string ->[]byte/或者[]rune ->修改 ->重写转成string
	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println(str)

	arr2 := []rune(str)
	arr2[0] = '猪'
	str = string(arr2)
	fmt.Println(str)
}
