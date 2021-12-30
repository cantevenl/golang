package main

import "fmt"

type MethpdUtils struct {
}

func (mu *MethpdUtils) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu *MethpdUtils) Print1(n, m int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu *MethpdUtils) Ares(len, width float64) float64 {
	return len * width
}

func (mu *MethpdUtils) Jiou(n int) {
	if n%2 == 0 {
		fmt.Println(n, "是偶数")
	} else {
		fmt.Println(n, "是奇数")
	}
}

func main() {
	var mu MethpdUtils
	mu.Print()
	mu.Print1(5, 3)
	res := mu.Ares(5, 9)
	fmt.Println(res)

	(&mu).Jiou(811)
}
