package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A里面的SayOK", a.Name)
}

func (a *A) hello() {
	fmt.Println("A里面的hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOk() {
	fmt.Println("B里面的SayOK", b.Name)

}

func main() {
	//var b B
	//b.A.Name = "猫猫"
	//b.A.age = 19
	//b.A.SayOk()
	//b.A.hello()
	//
	//var c B
	//c.Name="猪猪"
	//c.age=2
	//c.SayOk()
	//c.hello()
	//
	var b B
	b.Name = "猫猫"
	b.age = 100
	b.SayOk()
	b.hello()

	//如果非要给A结构体Name赋值
	b.A.Name = "牛牛"
	b.hello()
}
