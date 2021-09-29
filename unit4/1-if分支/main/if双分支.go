package main

import "fmt"

func main() {
	//实现功能：如果口罩小于50，提示不足，否正提示库存充足
	if count := 88; count < 50 {
		fmt.Println("库存不足")
	} else {
		fmt.Println("库存充足")
	}
}
