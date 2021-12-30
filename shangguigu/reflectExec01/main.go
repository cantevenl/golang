package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println(rType, rType.Kind())

	rVal := reflect.ValueOf(b)
	fmt.Println(rVal, rVal.Kind())

	iV := rVal.Interface()
	n := iV.(float64)
	fmt.Println(n)
}

func main() {
	var v float64 = 2.7
	reflect01(v)
}
