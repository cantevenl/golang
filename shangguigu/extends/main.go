package main

import "fmt"

//编写一个学生考试系统
//匿名结构体
type Student struct {
	Name  string
	Age   int
	Score float64
}

//将Pupil和Graduate共有的方法也绑定到 *Student
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩为%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score float64) {
	//业务判断
	if score >= 0 || score <= 100 {
		stu.Score = score
	} else {
		fmt.Println("获得的成绩不正确")
		return
	}
}

//给 *Student增加一个方法，那么Pupil和Graduate都可以使用该方法

func (stu *Student) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

//小学生
type Pupil struct {
	Student //嵌入了Student匿名结构体
}

//这时候Pupil结构体特有的方法，保留
func (p *Pupil) Testing() {
	fmt.Printf("小学生%v正在考试\n", p.Student.Name)
}

//大学生
type Graduate struct {
	Student
}

func (g *Graduate) Testing() {
	fmt.Printf("大学生%v正在考试\n", g.Name)
}

func main() {
	//当我们对结构体嵌入了匿名结构体使用方法方法会发现变化
	pupil := &Pupil{}
	pupil.Student.Name = "猪猪"
	pupil.Student.Age = 8
	pupil.Testing()
	pupil.Student.SetScore(99.5)
	pupil.Student.ShowInfo()
	fmt.Println(pupil.Student.GetSum(10, 20))

	graduate := &Graduate{}
	graduate.Student.Name = "孙子"
	graduate.Student.Age = 28
	graduate.Testing()
	graduate.Student.SetScore(100)
	graduate.Student.ShowInfo()
	fmt.Println(pupil.Student.GetSum(9, 27))
}
