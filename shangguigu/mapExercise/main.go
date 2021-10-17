package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {
	//判断users中是否有name
	//v,ok:=users[name]
	if users[name] != nil {
		//有这个用户
		users[name]["pwd"] = "888888"
	} else {
		//没有这个用户
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称~" + name
	}
}

func main() {
	users := make(map[string]map[string]string, 10)
	users["狗"] = make(map[string]string, 2)
	users["狗"]["pwd"] = "123456"
	users["狗"]["nickname"] = "狗狗"
	modifyUser(users, "猪")
	modifyUser(users, "猫")
	modifyUser(users, "狗")

	fmt.Println(users)
}
