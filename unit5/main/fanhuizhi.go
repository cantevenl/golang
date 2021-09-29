//package main
//
//import "fmt"
//
//func cal (n1 int,n2 int) {
//	//var sum int =0
//	sum := n1+n2
//	fmt.Println(sum)
//}
//
//func main() {
//	cal(27,25)
//}

//package main
//
//import "fmt"
//
//func cal (n1 int,n2 int) (int) {
//	//var sum int =0
//	sum := n1+n2
//	return sum
//}
//
//func main() {
//	sum := cal(10,20)
//	fmt.Println(sum)
//}

package main

import "fmt"

func niu(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	result := n1 - n2
	return sum, result
}

func main() {
	sum, result := niu(27, 25)
	fmt.Println(sum)
	fmt.Println(result)

}
