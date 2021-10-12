package main

import (
	"fmt"
	"time"
)

func main() {
	var user, passwd string
	var logintime = 3
AK:
	for {
		for ; logintime > 0; logintime-- {
			fmt.Println("请输入用户名和密码：")
			fmt.Scanf("%s%s", &user, &passwd)
			if user == "zhu" && passwd == "525" {
				fmt.Println("登陆成功\n")
				break AK
			} else {
				fmt.Printf("你还有%d次机会\n", logintime-1)
			}
			if logintime == 1 {
				fmt.Println("机会用完，请30分之后重试")
			}
		}
		time.Sleep(time.Second * 10)
		logintime = 3
	}

}
