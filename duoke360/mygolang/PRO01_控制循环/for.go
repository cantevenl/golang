package main

import "fmt"

func f5() {
	var a = [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Printf("i: %d, v: %v\n", i, v)
	}
}

func f6() {
	var s = []int{11, 22, 33, 44, 55}
	for i, v := range s {
		fmt.Printf("i:%d,v:%v\n", i, v)
	}
}

func f7() {
	m := make(map[string]string)
	m["name"]="niu"
	m["age"]="1"
	m["gender"]="man"

	for k,v := range m {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
}

func f8() {
	MYLABEL:
		for i:=0;i<=10;i++ {
			if i==5 {
				break MYLABEL
			}
			fmt.Println(i)
		}
		fmt.Println("end")
}

func main() {
	var num int
	for i := 0; i <= 100; i++ {
		num += i
	}
	fmt.Println(num)

	a := 1
	for a <= 10 {
		fmt.Println(a)
		a++
	}

	f5()
	f6()
	f7()
	f8()
}
