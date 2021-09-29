package main

import "fmt"

//定义一个猪居民的结构体，将居民中的各个属性 统一放入结构体中管理
type Zhu struct {
	//变量名字大写外界可以访问这个属性
	Name   string
	Age    int
	Gender string
}

func main() {
	//创建猪居民结构体的实例、对象、变量
	//猪国王；姓名：猪儿粑 年龄：2岁 性别：男
	var zhuerba Zhu
	fmt.Println(zhuerba) //在未赋值时默认值：{0}
	zhuerba.Name = "猪儿粑"
	zhuerba.Age = 2
	zhuerba.Gender = "男"
	fmt.Println(zhuerba)
	fmt.Println("十年后的年龄：", zhuerba.Age+10)

	//小老几
	var miaomiao Zhu
	miaomiao.Name = "小老几"
	miaomiao.Age = 1
	miaomiao.Gender = "女"
	fmt.Println(miaomiao)
}
