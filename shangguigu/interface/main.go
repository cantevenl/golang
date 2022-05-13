package main

import "fmt"

//声明(定义)一个接口
type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}

type Usb2 interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

type Camera struct {
}

//让Phone结构体实现Usb接口方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

//让Camera也实现Usb接口方法
func (p Camera) Start() {
	fmt.Println("相机开始工作")
}

func (p Camera) Stop() {
	fmt.Println("相机停止工作")
}

//计算机
type Computer struct {
}

//编写一个方法Working，方法接收一个Usb接口类型的变量
//只要是实现了Usb接口 （所谓实现Usb接口，就是指实现了Usb接口声明的所有方法）
func (c *Computer) Working(usb Usb) {
	//通过Usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

func main() {
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)
}
