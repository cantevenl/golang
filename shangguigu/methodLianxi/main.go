package main

import "fmt"

type MethodUtils struct {
	//字段...
}

//给MethodUtils编写方法
func (mu MethodUtils) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//m*n的矩形
func (mu MethodUtils) Print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//算矩形面积
func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

func main() {
	var mu MethodUtils
	mu.Print()
	fmt.Println()
	mu.Print2(3, 20)

	areaRes := mu.area(2.7, 3.9)
	fmt.Println()
	fmt.Println("面积为", areaRes)
}
