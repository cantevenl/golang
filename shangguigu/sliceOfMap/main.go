package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Weapon string
}

func main() {
	//1.声明一个map切片
	var mounsters []map[string]string
	mounsters = make([]map[string]string, 2)
	if mounsters[0] == nil {
		mounsters[0] = make(map[string]string)
		mounsters[0]["name"] = "孙悟空"
		mounsters[0]["age"] = "500"
		mounsters[0]["武器"] = "金箍棒"
	}
	if mounsters[1] == nil {
		mounsters[1] = make(map[string]string)
		mounsters[1]["name"] = "猪八戒"
		mounsters[1]["age"] = "400"
		mounsters[1]["武器"] = "耙子"
	}
	//下面这个写法不正确
	//if mounsters[2] ==nil{
	//	mounsters[2] =make(map[string]string)
	//	mounsters[2]["name"] = "沙僧"
	//	mounsters[2]["age"] = "300"
	//	mounsters[2]["武器"]="锄头"
	//}

	//使用切片的append函数，可以动态增加
	newmounster := map[string]string{
		"name": "沙僧",
		"age":  "300",
		"武器":   "锄头",
	}
	newmounster2 := map[string]string{
		"name": "唐僧",
		"age":  "200",
		"武器":   "无",
	}
	mounsters = append(mounsters, newmounster, newmounster2)

	for i, v := range mounsters {
		for k, v1 := range v {
			fmt.Printf("第%v个妖怪的%v是%v\n", i, k, v1)
		}
	}

	students := make(map[string]Student)
	student1 := Student{"孙悟空", 500, "金箍棒"}
	student2 := Student{"猪八戒", 400, "帕丁"}
	student3 := Student{"沙僧", 300, "锄头"}
	student4 := Student{"唐三藏", 200, "经书"}

	students["no1"] = student1
	students["no2"] = student2
	students["no3"] = student3
	students["no4"] = student4

	for k, v := range students {
		fmt.Printf("学生的编号是%v\n", k)
		fmt.Printf("学生的名字是%v\n", v.Name)
		fmt.Printf("学生的年龄是%v\n", v.Age)
		fmt.Printf("学生的地址是%v\n", v.Weapon)
		fmt.Println()
	}

}
