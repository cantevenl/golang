package main

import "fmt"

type Student struct {
	Name   string
	Gender string
	Age    int
	Id     int64
	Score  float64
}

type Box struct {
	len    float64
	width  float64
	height float64
}

type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) showPrice() {
	if visitor.Age >= 90 || visitor.Age <= 8 {
		fmt.Println("考虑到安全，不要玩耍")
		return
	}
	if visitor.Age > 18 {
		fmt.Printf("游客的名字为%v 年龄为%v 收费20元\n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("游客的名字为%v 年龄为%v 免费 \n", visitor.Name, visitor.Age)

	}
}

func (student Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]",
		student.Name, student.Gender, student.Age, student.Id, student.Score,
	)
	return infoStr
}

func (box Box) volume() float64 {
	return box.height * box.len * box.width
}

func main() {
	var student = Student{
		"猪猪",
		"男",
		2,
		888,
		99.99,
	}
	fmt.Println(student.say())

	var box = Box{
		2.2,
		3.3,
		4.4,
	}
	fmt.Printf("体积=%.2f", box.volume())

	var v Visitor
	for {
		fmt.Println("请输入你的名字：")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序....")
			break
		}
		fmt.Println("请输入你的年龄：")
		fmt.Scanln(&v.Age)
		v.showPrice()

	}
}
