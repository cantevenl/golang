package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int
	Num2 int
}

func (c *Cal) GetSub(name string) {
	res := c.Num1 - c.Num2
	fmt.Printf("%v完成了减法运算, %v-%v=%v\n", name, c.Num1, c.Num2, res)
}

func ReflectDemo(cal interface{}) {
	rValue := reflect.ValueOf(cal)
	rValueEle := rValue.Elem()
	num := rValueEle.NumField()

	for i := 0; i < num; i++ {
		fmt.Println(rValueEle.Field(i).Interface())
	}

	rMethod := rValue.Method(0)
	params := make([]reflect.Value, 0)
	params = append(params, reflect.ValueOf("Tom"))
	rMethod.Call(params)
}

func main() {
	c := Cal{
		Num1: 27,
		Num2: 9,
	}

	ReflectDemo(&c)
}
