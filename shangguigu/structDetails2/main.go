package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}
type B struct {
	Num int
}

type Student struct {
	Name string
	Age  int
}

type Stu Student

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A
	var b B
	a = A(b) //可以转换 但是要求有完全相同的字段
	fmt.Println(a, b)

	var stu1 Student
	var stu2 Stu
	//stu2 = stu1 不能这样
	stu2 = Stu(stu1)
	fmt.Println(stu1, stu2)

	//1.创建一个Monster变量
	monster := Monster{"牛魔王", 800, "芭蕉扇"}

	//2.将monster变量序列化为json格式字串
	//  json.Marshal 函数中使用反射
	jsonStr, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误", err)
	}
	fmt.Println("jsonStr", string(jsonStr))

}
