package main

import "fmt"

type MethodUtils struct {
}

func (mu *MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}

}

func (mu *MethodUtils) Print(n int, m int, key string) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

func main() {
	var mu MethodUtils
	fmt.Println("请输入数：")
	var num int
	fmt.Scanln(&num)
	mu.JudgeNum(num)

	mu.Print(9, 27, "#")

}
