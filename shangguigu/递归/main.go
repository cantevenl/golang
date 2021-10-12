package main

import "fmt"

//func test(n int) {
//	if n > 2 {
//		n--
//		test(n)
//	}
//	fmt.Println("n=", n)
//}

//func test(n int) {
//	if n > 2 {
//		n--
//		test(n)
//	} else {
//		fmt.Println("n=", n)
//	}
//}

//斐波那契
func feibo(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return feibo(n-1) + feibo(n-2)
	}
}

func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

func houzi(n int) int {
	if n == 10 {
		return 1
	} else {
		return (houzi(n+1) + 1) * 2
	}
}

func main() {
	res := feibo(8)
	fmt.Println(res)

	fmt.Println(f(2))

	fmt.Printf("第一共有%v个桃子", houzi(1))
}
