package main

import "fmt"

func main() {
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
