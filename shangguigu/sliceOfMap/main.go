package main

import "fmt"

func main() {
	//1.声明一个map切片
	var monsters []map[string]string
	monsters = make([]map[string]string, 2)
	//2.增加第一个妖怪的信息
	if monsters[0] == nil {
		monsters[0] = make(map[string]string, 2)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string, 2)
		monsters[1]["name"] = "狐狸精"
		monsters[1]["age"] = "100"
	}

	//这里我们需要使用到切片的append函数，动态增加一个map
	//1.先定义一个monster信息
	newMonster := map[string]string{
		"name": "新的妖怪",
		"age":  "200",
	}
	monsters = append(monsters, newMonster)

	fmt.Println(monsters)
}
