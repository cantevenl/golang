package main

import "fmt"

func modify(map1 map[int]int) {
	map1[10] = 999
}

//定义一个学生结构体
type Stu struct {
	Name string
	Age  int
	Addr string
}

func main() {
	//map是引用类型，在一个函数接收map，修改后，会直接修改原来的map

	map1 := make(map[int]int, 2)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	modify(map1)
	//看看结果 map1[10] = 999，说明map是引用类型
	fmt.Println(map1)

	//map的value经常使用struct类型，更适合管理复杂的数据
	//map的key为学生的学号，唯一的
	//map的value是结构体，包含学生的名字，年龄，地址

	students := make(map[string]Stu, 10)
	//创建学生
	stu1 := Stu{"猪", 18, "上海"}
	stu2 := Stu{"喵", 2, "成都"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

	//遍历各个学生的信息
	for k, v := range students {
		fmt.Printf("学生的编号是%v\n", k)
		fmt.Printf("学生的名字是%v\n", v.Name)
		fmt.Printf("学生的年龄是%v\n", v.Age)
		fmt.Printf("学生的地址是%v\n", v.Addr)
		fmt.Println()

	}
}
