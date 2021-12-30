package main

import "fmt"

type A struct {
	Name string
	Age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
}

type D struct {
	a A //有名结构体
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

type Monster struct {
	Name string
	Age  int
}

type E struct {
	Monster
	int
}

func main() {
	var c C
	//如果c 没有Name字段，而A和B都有相同的Name字段，这时候就必须通过指定匿名结构体名字来区分
	//所有c.Name就会报编译错误，这个规则对方法也是一样的
	c.A.Name = "牛牛"
	fmt.Println(c)

	var d D
	//如果是有名结构体，必须加名字
	d.a.Name = "猪猪"
	fmt.Println(d)

	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定各个匿名结构体字段的值
	tv := TV{Goods{"电视机", 5000.99}, Brand{"海尔", "山东青岛"}}

	//演示访问Goods的Name
	//fmt.Println(tv.Name)
	fmt.Println(tv.Goods.Name)

	tv2 := TV{
		Goods{
			"电视机2",
			9999.9,
		},
		Brand{
			"长虹",
			"成都",
		},
	}

	//使用指针
	tv3 := TV2{&Goods{"电视机3", 124214.99}, &Brand{"夏普", "北京"}}

	tv4 := TV2{
		&Goods{
			"电视机4",
			643564.9,
		},
		&Brand{
			Address: "日本",
			Name:    "索尼",
		},
	}

	fmt.Println("tv=", tv)
	fmt.Println("tv2=", tv2)
	fmt.Println("tv3=", *tv3.Goods, *tv3.Brand)
	fmt.Println("tv4=", *tv4.Goods, *tv4.Brand)

	//演示一下匿名字段是基本数据类型的使用
	var e E
	e.Monster.Name = "狐狸精"
	e.Age = 300
	e.int = 20
	fmt.Println(e)

}
