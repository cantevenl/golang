package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score int
}

func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

//小学生
type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中	")
}

//大学生
type Graduate struct {
	Student
}

func (p *Graduate) testing() {
	fmt.Println("大学生正在考试中	")
}


func main() {
	//当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Name="tom"
	pupil.Age=8
	pupil.testing()
	pupil.SetScore(66)
	pupil.ShowInfo()

	graduate := &Graduate{
		Student{"jack",
		20,
		0,
		},
	}
	graduate.testing()
	graduate.SetScore(88)
	graduate.ShowInfo()
}
