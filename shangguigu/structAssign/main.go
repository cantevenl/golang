package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}

func main() {
	//方式1
	//在创建结构体变量时，就直接指定字段的值
	var stu1 = Stu{"小明", 19} //stu1---> 结构体数据[xxx,xxx]
	stu2 := Stu{"小红", 20}

	//在创建结构体变量时，把字段名和字段值写在一起，这种写法就不依赖字段的定义顺序
	var stu3 = Stu{
		Age:  3,
		Name: "猪",
	}
	stu4 := Stu{
		"猫",
		2,
	}

	fmt.Println(stu1, stu2, stu3, stu4)

	//方式2：返回结构体的指针类型
	var stu5 = &Stu{"小王", 29} //stu5 --> 地址 ---> 结构体数据[xxx,xxx]
	stu6 := &Stu{"小牛", 39}

	//在创建结构体指针变量时，把字段名和字段值写在一起，这种写法，就不依赖字段的定义顺序
	var stu7 = &Stu{
		Name: "小李",
		Age:  49,
	}

	stu8 := &Stu{
		Age:  22,
		Name: "小张",
	}
	fmt.Println(stu5, stu6, stu7, stu8)
	fmt.Println(*stu5, *stu6, *stu7, *stu8)

}
