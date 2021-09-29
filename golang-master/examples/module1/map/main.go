package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]string, 10)
	myMap["name"] = "猪儿粑"
	myFuncMap := map[string]func() int{
		"funcA": func() int {
			return 27
		},
	}
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	fmt.Println(f())
	value, exists := myMap["name"]
	if exists {
		println(value)
	}
	for k, v := range myMap {
		println(k, v)
	}
}
