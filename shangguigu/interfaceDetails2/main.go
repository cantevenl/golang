package main

import "fmt"

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface
	test03()
}

//如果需要实现AInterface,就需要将B,C的方法都实现

type Stu struct {
}

func (stu Stu) test01() {
	fmt.Println("test01")
}

func (stu Stu) test02() {

}

func (stu Stu) test03() {

}

type T interface {
}

func main() {
	var stu Stu
	var a AInterface = stu
	//var a AInterface=Stu{}
	a.test01()

	var t T = stu
	fmt.Printf("%v\n", t)

	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	t = num1
	fmt.Println(t2, t)
}
