package main

import "fmt"

//声明(定义)一个接口
type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	name string
}

//让Phone结构体实现Usb接口方法
func (p Phone) Start() {
	fmt.Println(p.name, "手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println(p.name, "手机停止工作")
}

func (p Phone) Call() {
	fmt.Println(p.name, "手机 打电话")
}

type Camera struct {
	name string
}

//让Camera也实现Usb接口方法
func (c Camera) Start() {
	fmt.Println(c.name, "相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println(c.name, "相机停止工作")
}

type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	//如果usb时指向Phone结构体变量，则还需要调用call方法
	//类型断言
	if phone, ok := usb.(Phone); ok == true {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	//定义一个Usb接口数组，可以存放Phone和Camear的结构体变量
	//这里就体现出多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"苹果"}
	usbArr[1] = Phone{"华为"}
	usbArr[2] = Camera{"索尼"}

	//遍历usbArr
	//Phone还有一个特有的方法call()，请遍历Usb数组，如果是Phone变量，
	//除了调用Usb 接口声明的方法外，还需要调用Phone特有方法call
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
	//fmt.Println(usbArr)
}
