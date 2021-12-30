package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"monster_name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

func TestSturct(a interface{}) {
	tye := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	//tye_kd:=tye.Kind()
	//fmt.Println(tye_kd)
	val_kd := val.Kind()
	fmt.Println(val_kd)
	if val_kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("白象精")
	val.Elem().Field(1).SetInt(888)
	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Elem().Field(i).Kind())
	}

	fmt.Printf("sturct has %d fields\n", num)

	tag := tye.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%v\n", tag)

	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d method\n", numOfMethod)

	val.Elem().Method(0).Call(nil)
}

func main() {
	m := Monster{"猪八戒", 999, 25.55, "男"}

	//Marshal就是通过反射获取到struct的tag值
	result, _ := json.Marshal(m)
	fmt.Println("json result：", string(result))

	TestSturct(&m)
	fmt.Println(m)

}
