package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

func main() {
	//定义一个 Usb 接口数组，可以存放 Phone 和 Camera 的结构体变量
	//这里体现多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"苹果"}
	usbArr[1] = Phone{"华为"}
	usbArr[2] = Camera{"索尼"}

	fmt.Println(usbArr)
}
