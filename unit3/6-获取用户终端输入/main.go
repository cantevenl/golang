//package main
//
//import "fmt"
//
//func main() {
//	//实现功能：键盘录入学生的年龄，姓名，成绩，是否是VIP
//	//方式一：Scanln
//	var age int
//	fmt.Println("请录入学生的年龄：")
//	//传入age的地址的目的：在Scanln函数中，对地址中的值进行改变的时候，实际外面的age被影响了
//	fmt.Scanln(&age)
//
//	var name string
//	fmt.Println("请输入学生的名字：")
//	fmt.Scanln(&name)
//
//	var score float32
//	fmt.Println("请输入学生的成绩：")
//	fmt.Scanln(&score)
//
//	var vip bool
//	fmt.Println("请输入学生是否是vip:")
//	fmt.Scanln(&vip)
//
//	//将上述数据在控制台打印出来
//	fmt.Printf("学生的年龄为：%v，姓名为：%v，成绩为：%v，是否是vip：%v", age, name, score, vip)
//
//}

package main

import "fmt"

func main() {
	var age int
	var name string
	var score float32
	var vip bool
	fmt.Printf("学生的年龄,姓名,成绩,是否是vip,使用空格进行分隔\n")
	fmt.Scanf("%d %s %f %t", &age, &name, &score, &vip)

	fmt.Printf("学生的年龄为：%v，姓名为：%v，成绩为：%v，是否是vip：%v", age, name, score, vip)
}
