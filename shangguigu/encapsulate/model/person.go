package model

import "fmt"

type person struct {
	Name string
	age  int //其他包不能直接访问
	sal  float64
}

type account struct {
	username string
	balance  float64
	passwd   string
}

func NewAccount(username string, balance float64, passwd string) *account {
	if len(username) < 6 || len(username) > 10 {
		fmt.Println("账号长度有误")
		return nil
	}

	if balance < 20 {
		fmt.Println("余额不对")
		return nil
	}

	if len(passwd) != 6 {
		fmt.Println("密码成都不对")
		return nil
	}

	return &account{
		username,
		balance,
		passwd,
	}
}

func (a *account) Setname(username string) {
	if 6 < len(username) || len(username) < 10 {
		a.username = username
	} else {
		fmt.Println("账户名输入错误")
	}
}

func (a *account) Setbalance(balance float64) {
	if balance > 20 {
		a.balance = balance
	} else {
		fmt.Println("余额输入有误")
	}
}

func (a *account) Setpasswd(passwd string) {
	if len(passwd) == 6 {
		a.passwd = passwd
	} else {
		fmt.Println("密码长度不对")
	}
}

func (a *account) Getname() string {
	return a.username
}

func (a *account) Getbalance() float64 {
	return a.balance
}

func (a *account) Getpassswd() string {
	return a.passwd
}

//写一个工厂模式的函数，相当于构造函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//为了访问age和sal我们编写一对Set方法和Get方法
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *person) Getage() int {
	return p.age
}

func (p *person) Setsal(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确")
	}
}

func (p *person) Getsal() float64 {
	return p.sal
}
