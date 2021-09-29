package main

import "fmt"

func main() {
	//逻辑与	只要有一个false就为false
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(false && true)

	//逻辑或 只要有一个true就为true
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(false || true)

	//逻辑非，取反的结果
	fmt.Println(!true)
	fmt.Println(!false)
}
