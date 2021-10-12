package main

import "fmt"

func main() {
	//1.Scanln
	var name string
	var age int
	var sal float32
	var pass bool
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过考试(true/false)：")
	fmt.Scanln(&pass)

	fmt.Printf("名字是：%v,年龄：%v,薪水:%v,是否通过考试:%v\n", name, age, sal, pass)
	//2.Scanf	按指定的格式输入
	fmt.Println("请输入姓名，年龄，薪水，是否通过考试，使用空格隔开")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &pass)
	fmt.Printf("名字是：%v,年龄：%v,薪水:%v,是否通过考试:%v", name, age, sal, pass)

}
