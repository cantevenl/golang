package main

import "fmt"

func main() {
	//输出1-100内能被6整除的数
	for i := 1; i <= 100; i++ {
		if i%6 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("方式二。。。。。。。。。。。。。。。。")
	//continue
	for i := 1; i <= 100; i++ {
		if i%6 != 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("加标签。。。。。。。。。。")
lable:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 4; j++ {
			if i == 2 && j == 2 {
				continue lable
			}
			fmt.Printf("i:%v,j:%v\n", i, j)
		}
	}
}
