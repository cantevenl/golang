package main

import (
	"fmt"
	"shangguigu/factory/model"
)

func main() {
	//创建一个Student实例
	//var stu = model.student{
	//	Name:  "tom",
	//	Score: 78.9,
	//}

	//当student结构体首字母小写，我们可以通过工厂模式来解决
	var stu = model.NewStudent("tom~", 88.8)
	fmt.Println(stu)
	fmt.Println(*stu)
	fmt.Println("name=", stu.Name, "score=", stu.GetScore())
}
