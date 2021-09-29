package main

import "fmt"

func main() {
	//根据给出的学生分数，判断学生的等级：
	var score float32
	fmt.Println("请输入你的成绩：")
	if fmt.Scanln(&score); score >= 90 {
		fmt.Println("你牛逼")
	} else if score >= 80 {
		fmt.Println("你还行")
	} else if score >= 70 {
		fmt.Println("你一般")
	} else if score >= 60 {
		fmt.Println("你垃圾")
	} else if score < 60 {
		fmt.Println("你是废物")
	}
}
