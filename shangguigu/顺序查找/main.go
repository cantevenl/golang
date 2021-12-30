package main

import "fmt"

func main() {
	//有一个数列：猪、牛、猫、狗
	//从键盘输入一个名称，判断数列中是否包含此名称（顺序查找）

	name := [4]string{"猪", "牛", "猫", "狗"}
	var animal string = ""
	fmt.Println("请输入要查找的动物：")
	fmt.Scanln(&animal)

	//顺序查找，第一种方式
	//for i := 0; i < len(name); i++ {
	//	if animal == name[i] {
	//		fmt.Printf("找到动物%v,下标为%v\n", name[i], i)
	//		break
	//	} else if i == (len(name) - 1) {
	//		fmt.Printf("没有找到%v\n", animal)
	//	}
	//}
	//顺序查找，第二种方式（推荐使用）
	index := -1
	for i := 0; i < len(name); i++ {
		if animal == name[i] {
			index = i //将找到的值对应的下标赋给index
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到动物%v,下标为%v", name[index], index)
	} else if index == -1 {
		fmt.Println("没找到")
	}

}
