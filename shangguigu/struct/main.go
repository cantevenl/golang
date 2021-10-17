package main

import "fmt"

//如果结构体的字段类型是：指针，slice，map 它们零值都是nil，即还没有分配空间
//如果需要使用这样的字段，需要先make，才能使用

type Person struct {
	Name   string
	Age    int
	scores [5]float64
	ptr    *string           //指针
	slice  []int             //切片
	map1   map[string]string //map
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	//定义一个结构体变量
	var p1 Person
	fmt.Println(p1)

	//使用slice
	p1.slice = make([]int, 10)
	p1.slice[0] = 27

	//使用map，一定要先make
	p1.map1 = make(map[string]string)
	p1.map1["爱好"] = "游戏,篮球"
	fmt.Println(p1)

	//使用ptr需要使用new
	a := "啊"
	p1.ptr = new(string)
	p1.ptr = &a
	fmt.Println(p1)

	//不同结构体变量的字段是独立的，互不影响，一个结构体变量字段的更改，
	//不影响另外一个。结构体是值类型
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 9999

	monster2 := monster1 //结构体是值类型，默认为值拷贝
	monster2.Name = "青牛精"

	fmt.Println(monster1) //{牛魔王 9999}
	fmt.Println(monster2) //{青牛精 9999}

	//但是如果结构体里面是引用类型，那就能够改变
	p2 := p1
	p2.map1["爱好"] = "睡觉"
	fmt.Println(p1)
	fmt.Println(p2)

}
