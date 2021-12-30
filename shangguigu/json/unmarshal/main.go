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

//将json字符串，反序列化成struct
func unmarshalStruct() {
	//这个str 在项目开发中，是通过网络传输获取到的，或者是读取文件获取的
	str := "{\"monster_name\":\"牛魔王\",\"monster_age\":500,\"monster_birthday\":\"2011-11-11\",\"monster_sal\":8000,\"monster_skill\":\"牛魔拳\"}"

	//定义一个Monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err =%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v\n monster.Name=%v\n", monster, monster.Name)
}

//将json字符串，反序列化成map
func unmarshalMap() {
	str := "{\"address\":\"火云洞\",\"age\":30,\"name\":\"红孩儿\"}\n"

	//定义一个Map
	var a map[string]interface{}

	//反序列化
	//注意：反序列化map，不需要make，因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err =%v\n", err)
	}
	fmt.Printf("反序列化后 a=%v\n", a)

}

//将json字符串，反序列化成slice
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"},{\"address\":[\"墨西哥\",\"夏威夷\",\"南极\"],\"age\":\"27\",\"name\":\"tom\"}]"

	//定义一个slice
	var slice []map[string]interface{}
	//反序列化
	//反序列化slice，不需要make，因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err =%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)

}

func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
