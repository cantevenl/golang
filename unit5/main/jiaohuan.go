package main

import "fmt"

//自定义函数，交换两个数
func exchange(n1 int, n2 int) {
	var t int
	t = n1
	n1 = n2
	n2 = t
	fmt.Println(n1,n2)
}

func main() {
	n1 := 27
	n2 := 9
	fmt.Printf("交换前的两个数：n1=%v,n2=%v\n", n1, n2)
	exchange(n1, n2)
	fmt.Printf("交换后的两个数：n1=%v,n2=%v\n", n1, n2)

}
