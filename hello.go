package main

import "fmt"

//func feibo(n int) int {
//	if n == 1 || n == 2 {
//		return 1
//	} else {
//		return feibo(n-1) + feibo(n-2)
//	}
//}
//
//func niu(n int) int{
//	if n==1 {
//		return 3
//	}else {
//		return 2*niu(n-1)+1
//	}
//}
//
//func main() {
//	var num int
//	fmt.Println("求第几个斐波那契数")
//	fmt.Scanln(&num)
//	fmt.Println(feibo(num))
//
//	fmt.Println(niu(5))
//}

func fbn(n int) uint64 {
	var arr []uint64
	arr = make([]uint64, 10)
	arr[0] = 1
	arr[1] = 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	fmt.Println(arr)
	return arr[n-1]
}

func main() {
	var n int
	fmt.Println("想求第几个斐波那契数：")
	fmt.Scanln(&n)
	res := fbn(n)
	fmt.Printf("第%v个斐波那契数是：%v", n, res)
}
