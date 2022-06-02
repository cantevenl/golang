package main

import "fmt"

func f9() {
	// MY_LABEL:
	for i := 0; i < 5; i++ {
	MY_LABEL:
		for j := 0; j < 5; j++ {
			if i == 4 && j == 4 {
				continue MY_LABEL
			}
			fmt.Printf("i=%d,j=%d\n", i, j)
		}
	}
}

func f10() {
	a:=0
	if a==1{
		goto NIU
	}else {
		goto ZHU
	}

	NIU:
		fmt.Println("牛")

	ZHU:
		fmt.Println("猪")
}

func f18() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				goto LABEL1
			}
		}
	}
LABEL1:
	fmt.Println("label1")
}

func main() {
	for i:=0;i<=100;i++ {
		if i%2==0 {
			fmt.Printf("i: %v\n", i)
		}
	}

	f9()
	f10()
	f18()
}
