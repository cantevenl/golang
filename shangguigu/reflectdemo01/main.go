package main

import (
	"fmt"
	"reflect"
)

//专门演示反射
func reflectTest01(b interface{}) {
	//通过反射获取到传入的变量的type，kind，值
	//1.先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

	n2 := 2 + rVal.Int()
	//n3 := rVal.Float()
	fmt.Println("n2=", n2)
	//fmt.Println("n3=", n3)

	fmt.Printf("rVal=%v\t rVal的type=%T\n", rVal, rVal)

	//3.将rVal转成interface{}
	iV := rVal.Interface()
	//将interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Printf("num2=%v\t num2的type=%T\n", num2, num2)
}

//对结构体的反射
func reflectTest02(b interface{}) {
	//1.先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("reflectTest02 rType=", rType)

	//2.获取到reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

	//3.获取变量对应的Kind
	//(1) rVal.Kind()
	kind1 := rVal.Kind()
	//(2) rTyp.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind=%v\tkind=%v\n", kind1, kind2)

	iV := rVal.Interface()
	fmt.Printf("iv=%v\t iv的type=%T\n", iV, iV)
	//将interface{}通过断言转成需要的类型
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v", stu.Name)
	}

}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	//对基本数据类型、interface{}、reflect.Value进行反射的基本操作

	//1.先定义一个int
	var num int = 100
	reflectTest01(num)

	//2.定义一个Student的实例
	stu := Student{"猪", 6}
	reflectTest02(stu)
}
