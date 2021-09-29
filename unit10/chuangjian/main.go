package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	var t *Student = new(Student)
	//t是指针，t其实指向的就是地址，应该给这个地址指向的对象的字段进行赋值操作
	(*t).Name = "猪儿粑"
	(*t).Age = 2

	//为了符合编程习惯，go提供了简化的赋值方式
	t.Gender = "男"
	//go编译器底层对t.Gender转化(*t).School = "男"
	fmt.Println(*t)

}
