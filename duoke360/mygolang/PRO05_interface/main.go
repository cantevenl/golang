package main

import "fmt"

type USB interface {
	Start()
	Stop()
}

type Phone struct {
}

func (p *Phone) Start() {
	fmt.Println("手机开始工作。。。\n")
}

func (p *Phone) Stop() {
	fmt.Println("手机停止工作。。。\n")
}

type Camera struct {
}

func (c *Camera) Start() {
	fmt.Println("相机开始工作。。。\n")
}

func (c *Camera) Stop() {
	fmt.Println("相机停止工作。。。\n")
}

type Computer struct {

}

func (com *Computer) Working(usb USB) {
	usb.Start()
	usb.Stop()
}

func main() {
	computer:=Computer{}
	phone:=Phone{}
	camera:=Camera{}

	computer.Working(&phone)
	computer.Working(&camera)

}
