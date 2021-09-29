package main

import "fmt"

func main() {
	var scores [5]int
	//将成绩存入数组：（循环+终端输入）
	for i := 0; i < len(scores); i++ { //i:数组的下标
		fmt.Printf("请录入第%d个学生的成绩：", i+1)
		fmt.Scanln(&scores[i])
	}

	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]

	}
	fmt.Println(sum)
	avg := sum / len(scores)
	fmt.Println(avg)

	//展示一些班级每个学生的成绩：（对数组进行遍历）
	//方式1：普通for循环：
	for i := 0; i < len(scores); i++ {
		fmt.Printf("第%d个学生的成绩为：%d\n", i+1, scores[i])
	}

	fmt.Println("-------------------------------------")

	//方式2：for range循环
	for k, v := range scores {
		fmt.Printf("第%d个学生的成绩为：%d\n", k, v)
	}

	fmt.Println("-------------------------------------")
	//如果想忽略某个值用_就可以
	for _, v := range scores {
		fmt.Printf("学生的成绩为：%d\n", v)
	}
}
