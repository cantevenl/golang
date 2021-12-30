package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	//获取reflect.Type类型
	typ := reflect.TypeOf(a)
	//获取到reflect.Value类型
	val := reflect.ValueOf(a)
	//获取到a对应的类别
	kd := val.Kind()
	//如果传入的不是结构体，就退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取到该结构体有几个字段
	num := val.NumField()
	fmt.Printf("结构体有%d个fields\n", num) //4
	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d：值为=%v\n", i, val.Field(i))
		//获取到struct标签，注意需要通过reflect.Type来获取tag标签的值
		tagVal := typ.Field(i).Tag.Get("json")
		//如果该字段有tag标签就显示，否则就不显示了
		if tagVal != "" {
			fmt.Printf("Field %d：tag为=%v\n\n", i, tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("结构体有%d个方法\n", numOfMethod)

	//排序是按照方法名字母排序（ASCII码）
	//获取到下标为1的方法（第二个方法）
	//然后调用该方法
	val.Method(1).Call(nil)

	//调用结构体的第一个方法Method(0)
	var params []reflect.Value //声明了[]reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Method(0).Call(params) //传入的参数是[]reflect.Value，返回的还是[]reflect.Value
	fmt.Println("res=", res[0].Int()) //返回结果，返回的结果是[]reflect.Value

}

func main() {
	var a Monster = Monster{
		"白骨精",
		400,
		30.8,
		"女",
	}
	//将Monster实例传递给TestStruct
	TestStruct(a)

}
