package model

import "fmt"

//声明一个Customer结构体，表示一个客户信息

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

//使用工厂模式，返回一个*customer的实例
func Newcustomer(id int, name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{
		id,
		name,
		gender,
		age,
		phone,
		email,
	}
}

//第二种创建Customer实例方法，不带id
func Newcustomer2(name string, gender string, age int, phone string, email string) *Customer {
	return &Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

//返回用户的信息
func (this Customer) Getinfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t", this.Id, this.Name, this.Gender,
		this.Age, this.Phone, this.Email)
	return info
}

//返回用户的单个信息
func (this Customer) GetOneInfo(info string) string {
	switch info {
	case "Name":
		return this.Name
	case "Gender":
		return this.Gender
	case "Age":
		return string(this.Age)
	case "Phone":
		return this.Phone
	case "Email":
		return this.Email
	default:
		return ""
	}
}
