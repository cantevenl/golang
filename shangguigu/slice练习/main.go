package main

import "fmt"

func fbn(n int) []uint64 {
	//声明一个切片，切片大小 n
	fbnslice := make([]uint64, n)
	//第一，二个数都为1
	fbnslice[0] = 1
	fbnslice[1] = 1
	for i := 2; i < n; i++ {
		fbnslice[i] = fbnslice[i-1] + fbnslice[i-2]
	}
	return fbnslice
}

func main() {
	fnbSlice := fbn(20)
	fmt.Println(fnbSlice)
}
