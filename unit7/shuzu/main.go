package main

import "fmt"

func main() {
	//给出五个学生的成绩，求出成绩的总和以及平均数
	//数组存储相同类型的数据
	//定义一个数组
	var scores [5]int
	//将成绩存入数组：
	scores[0] = 97
	scores[1] = 27
	scores[2] = 58
	scores[3] = 100
	scores[4] = 86
	//求和：
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}
	avg := sum / len(scores)

	fmt.Printf("成绩总和：%v,成绩平均值：%v", sum, avg)

}
