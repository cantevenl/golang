package main

import (
	"encoding/json"
	"fmt"
)

//定义一个结构体
type Monster struct {
	//sturct的每个字段上，可以写上一个tag，该tag可以通过==反射机制==获取，常见的使用场景就是序列化和反序列化
	Name     string  `json:"monster_name"`
	Age      int     `json:"monster_age"`
	Birthday string  `json:"monster_birthday"`
	Sal      float64 `json:"monster_sal"`
	Skill    string  `json:"monster_skill"`
}

//将结构体序列化
func testStruct() {
	//演示
	monster := Monster{
		"牛魔王",
		500,
		"2011-11-11",
		8000.0,
		"牛魔拳",
	}

	//将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误 err=", err)
	}
	//输出序列化后的结果
	fmt.Printf("序列化后的结果=%v\n", string(data))
}

//将map进行序列化
func testMap() {
	//map里的key是字符串，值是任意类型（空接口代表任意类型）
	var a map[string]interface{}
	//使用map之前需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"

	//将a这个map序列化
	data, err := json.Marshal(a) //这里不需要传至针，因为map本身就是引用类型
	if err != nil {
		fmt.Println("序列化错误 err=", err)
	}
	fmt.Printf("a map 序列化后的结果=%v\n", string(data))
}

//对切片进行序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	//使用map前要make
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	//使用map前要make
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "27"
	m2["address"] = [3]string{"墨西哥", "夏威夷", "南极"}
	slice = append(slice, m2)

	//将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化错误 err=", err)
	}
	fmt.Printf("slice 序列化后=%v\n", string(data))

}

//对基础数据类型进行序列化，对基本数据序列化意义不大
func testFloat64() {
	var num1 float64 = 2345.67
	//将num1进行序列化操作
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("序列化错误 err=", err)
	}
	fmt.Printf("num1 序列化后=%v\n", string(data))
}

func main() {
	//将结构体、map、切片进行序列化
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}
