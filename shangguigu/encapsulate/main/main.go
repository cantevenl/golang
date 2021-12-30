package main

import (
	"fmt"
	"golang/shangguigu/encapsulate/model"
)

func main() {
	var p = model.NewPerson("猪猪")
	//p:=model.NewPerson()
	p.SetAge(18)
	p.Setsal(8000)
	fmt.Println(p)

	fmt.Printf("%v的年龄是%v，目前薪水%v\n", p.Name, p.Getage(), p.Getsal())

	//第二题
	var username string
	var passwd string
	var balance float64
	fmt.Println("请输入你的名字")
	fmt.Scanln(&username)
	fmt.Println("请输入你的密码")
	fmt.Scanln(&passwd)
	fmt.Println("请输入你的余额")
	fmt.Scanln(&balance)
	a := model.NewAccount(
		username,
		balance,
		passwd,
	)
	if a != nil {
		fmt.Println("创建成功=", a)
	} else {
		fmt.Println("创建失败")
	}
	a.Setpasswd("Aa123!")

	fmt.Printf("%v的密码是%v，目前薪水%v美元\n", a.Getname(), a.Getpassswd(), a.Getbalance())
}
