package main

import "fmt"

type Monkey struct {
	Name string
}

//声明接口
type BirdAble interface {
	Flying()
}

type FishAble interface {
	Swimming()
}

func (monkey *Monkey) climbing() {
	fmt.Println(monkey.Name, "生来会爬树")
}

//LittleMonkey结构体
type LittleMonkey struct {
	Monkey //继承
}

//让LittleMonkey实现BirdAble
func (monkey *LittleMonkey) Flying() {
	fmt.Println(monkey.Name, "通过学习，会飞翔。。。")
}

//实现FishAble
func (monkey *LittleMonkey) Swimming() {
	fmt.Println(monkey.Name, "通过洗澡，学会游泳")
}

func main() {
	//创建一个LittleMonkey 实例
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}

	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()
}
