package main

import "fmt"

type AInterface interface {
	Say()
}

type BInterface interface {
	Hello()
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster里Hello()")
}

func (m Monster) Say() {
	fmt.Println("Monster里Say()")
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i = ", i)
}

func main() {
	var stu Stu //结构体变量，实现了Say()方法 实现了 AInterface
	var a AInterface = stu
	a.Say()

	var i integer = 109
	var b AInterface = i
	b.Say()

	//Monster实现了AInterface和BInterface两个接口
	var m Monster
	var a1 AInterface = m
	var b1 BInterface = m
	a1.Say()
	b1.Hello()

}
