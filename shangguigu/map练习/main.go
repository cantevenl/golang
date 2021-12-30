package main

import "fmt"

/*
1.使用map[string]map[string]string 的map类型

2.key:表示用户名，是唯一的，不可以重复

3.如果某个用户名存在，就将其密码修改为"888888"，如果不存在就增加这个用户信息（包括昵称nickname和密码pwd）

4.编写一个函数modifyUser(users map[string]map[string]string, name string)
*/

func modifyUser(users map[string]map[string]string, name string, nickname string) {
	//判断users里是否有name
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		//没有没有用户
		users[name] = make(map[string]string)
		users[name]["nickname"] = nickname
		users[name]["pwd"] = "888888"
	}
}

func main() {
	var users map[string]map[string]string
	users = make(map[string]map[string]string)
	users["猪儿粑"] = map[string]string{
		"nickname": "猪国王",
		"pwd":      "123456",
	}
	modifyUser(users, "猪儿粑", "猪子")
	modifyUser(users, "小老几", "迪士尼老大")
	modifyUser(users, "圣诞儿", "最小的老几")

	for k, v := range users {
		for k1, v1 := range v {
			fmt.Printf("%v的%v是%v\n", k, k1, v1)
		}
	}
}
