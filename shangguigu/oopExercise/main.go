package main

import "fmt"

type Student struct {
	Name   string
	Gender string
	Age    int
	Id     int
	Score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		student.Name, student.Gender, student.Age, student.Id, student.Score,
	)
	return infoStr
}

func main() {
	//创建一个Studnet实例变量
	var stu = Student{
		"猪",
		"男",
		8,
		1000,
		99.8,
	}
	fmt.Println(stu.say())
}
