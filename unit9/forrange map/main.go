package main

import "fmt"

func main() {
	b := make(map[int]string)
	b[110] = "牛"
	b[220] = "猪"
	b[330] = "猫"

	//获取长度
	fmt.Println(len(b))

	//遍历
	for k, v := range b {
		fmt.Printf("key:%v,value:%v\n", k, v)
	}

	//加深难度
	a := make(map[string]map[int]string)
	//赋值
	a["班级1"] = make(map[int]string, 3)
	a["班级1"][110] = "猪"
	a["班级1"][120] = "牛"
	a["班级1"][130] = "鱼"

	a["班级2"] = make(map[int]string, 3)
	a["班级2"][111] = "虎"
	a["班级2"][122] = "狮"
	a["班级2"][133] = "豹"

	for k1, v1 := range a {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Printf("学生学号为：%v，学生姓名为：%v\t", k2, v2)

		}
		fmt.Println()
	}
}
