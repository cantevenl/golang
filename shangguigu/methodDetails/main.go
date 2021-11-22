package main

import "fmt"

//自定义类型，都可以有方法

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

//编写一个方法，可以改变i的值
func (i *integer) change() {
	*i += 1
}

//如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个变量的String()进行输出
type Student struct {
	Name string
	Age  int
}

//给Student实现方法String()
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func main() {
	var i integer = 10
	i.print()
	i.change()
	fmt.Println("i=", i)

	//定义一个Student变量
	stu := Student{
		"tom",
		19,
	}
	//如果你实现了 *Student 类型的 String方法，就会自动调用
	fmt.Println(&stu)

}
