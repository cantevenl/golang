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
	name string
}

//让Phone结构体实现Usb接口方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

type Camera struct {
	name string
}

//让Camera也实现Usb接口方法
func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

func main() {
	//定义一个Usb接口数组，可以存放Phone和Camear的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"苹果"}
	usbArr[1] = Phone{"华为"}
	usbArr[2] = Camera{"索尼"}
	fmt.Println(usbArr)
}
