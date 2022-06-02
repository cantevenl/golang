package main

import "fmt"

func f(score int) {
	if score >=60 && score <=70 {
		fmt.Println("C")
	}else if score > 70 && score <=90 {
		fmt.Println("B")
	}else if score > 90{
		fmt.Println("A")
	}else {
		fmt.Println("垃圾")
	}
}

func f1() {
	a:=10
	b:=20
	c:=3
	if a>b{
		if a>c {
			fmt.Println("a最大")
		}
	}else {
		if b >c {
			fmt.Println("b最大")
		}else {
			fmt.Println("c最大")
		}
	}
}


func main() {
	var score int
	fmt.Println("输入分数：")
	fmt.Scanln(&score)

	f(score)
	f1()
}
