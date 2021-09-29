package main

import "fmt"

func main() {
	a := make(map[int]string)
	//增加
	a[110] = "牛"
	a[120] = "猪"
	//修改
	a[110] = "猪警察"
	fmt.Println(a)

	//删除
	delete(a, 120)
	fmt.Println(a)

	//清空
	a = make(map[int]string)
	a[1] = "牛"
	a[2] = "猪"
	fmt.Println(a)

	//查找
	value, flag := a[1]
	fmt.Println(value)
	fmt.Println(flag)
	value1, flag1 := a[124]
	fmt.Println(value1)
	fmt.Println(flag1)
}
