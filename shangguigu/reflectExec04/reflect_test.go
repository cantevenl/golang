package main

import (
	"fmt"
	"reflect"
	"testing"
)

type GMethod struct {
	Func interface{}
	Args []interface{}
}

func test1(a, b int) {
	fmt.Println(a, b)
}

func test2(a, b int, str string) {
	fmt.Println(a, b, str)
}

func sum(a, b int) int {
	return a + b
}

func Run(method GMethod) {
	_func := method.Func
	args := method.Args

	rFunc := reflect.ValueOf(_func)
	rArgs := make([]reflect.Value, 0)
	for _, arg := range args {
		rArgs = append(rArgs, reflect.ValueOf(arg))
	}

	rFunc.Call(rArgs)
}

func RunWithReturn(method GMethod) []interface{} {
	_func := method.Func
	args := method.Args

	rFunc := reflect.ValueOf(_func)
	rArgs := make([]reflect.Value, 0)
	for _, arg := range args {
		rArgs = append(rArgs, reflect.ValueOf(arg))
	}

	rRet := rFunc.Call(rArgs)
	actRet := make([]interface{}, 0)
	for _, ele := range rRet {
		actRet = append(actRet, ele.Interface())
	}

	return actRet
}

func TestMultiFunc(t *testing.T) {
	func1 := GMethod{
		Func: test1,
		Args: []interface{}{1, 2},
	}

	Run(func1)

	func2 := GMethod{
		Func: test2,
		Args: []interface{}{1, 2, "Hello"},
	}

	Run(func2)
}

func TestMultiFuncWithReturn(t *testing.T) {
	func3 := GMethod{
		Func: sum,
		Args: []interface{}{10, 20},
	}

	ret := RunWithReturn(func3)
	fmt.Println(ret[0])
}
