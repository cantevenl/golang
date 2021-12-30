package main

import "fmt"

func main() {

	var money interface{} = "Niu"
	switch money.(type) {
	case int:
		fmt.Println("is int")
	case int64:
		fmt.Println("is int64")
	case float64:
		fmt.Println("is float64")
	case float32:
		fmt.Println("is float32")
	case string:
		fmt.Println("is string")
	default:
		fmt.Println("未知")
	}

}
