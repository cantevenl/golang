package main

import "fmt"

type Person struct {
	name, addr string
	id, age    int
}

func (per Person) eat() {
	fmt.Println(per.name + " eating....")
}
func (per Person) sleep() {
	fmt.Println(per.name + " sleep....")
}

func showPerson(p Person) {
	p.id = 1
	p.age = 88
	p.name = "kite"
	p.addr = "shanghai"
}

func showPerson1(p *Person) {
	p.id = 1
	p.age = 88
	p.name = "kite"
	p.addr = "shanghai"
}

type Dog struct {
	name  string
	color string
	age   int
}

type Master struct {
	dog  Dog
	name string
	age  int
}

type Ren struct {
	name string
}

func showRen(r Ren) {
	fmt.Println("showRen",r.name)
	r.name="特朗普"
	fmt.Println("showRen",r.name)
}

func showRen1(r *Ren) {
	fmt.Println("showRen1",r.name)
	r.name="特朗普"
	fmt.Println("showRen1",r.name)
}

func (r Ren) showRen(){
	fmt.Println("showRen",r.name)
	r.name="特朗普"
	fmt.Println("showRen",r.name)
}

func (r *Ren) showRen1(){
	fmt.Println("showRen1",r.name)
	r.name="特朗普"
	fmt.Println("showRen1",r.name)
}


func main() {
	//p:=Person{
	//	"tom",
	//	18,
	//}
	//fmt.Println(p.name)

	//匿名结构体
	var d struct {
		id   int
		name string
	}
	d.id = 1
	d.name = "牛"
	fmt.Println(d)

	p := Person{
		"tom",
		"chengdu",
		1,
		20,
	}
	fmt.Println(p, p.age)

	pp := &Person{}
	fmt.Printf("%T\n,%v\n,%v", pp, pp, *pp)
	pp.name = "zhu"
	pp.addr = "beijing"
	pp.id = 100
	pp.age = 2
	fmt.Printf("%T\n,%v\n,%v", pp, pp, *pp)

	var p_person = new(Person)
	fmt.Printf("p_person: \n%T\n", p_person)
	p_person.id = 99
	p_person.name = "gou"
	fmt.Printf("*p_person: %v\n", *p_person)

	ppp := new(Person)
	fmt.Println(ppp)

	person := Person{
		"zhangsan",
		"riben",
		50,
		50,
	}
	fmt.Printf("person: %v\n", person)
	fmt.Println("----------------")
	showPerson(person)
	fmt.Println("----------------")
	fmt.Printf("person: %v\n", person)
	fmt.Println("-------指针传递，地址传递---------")
	showPerson1(&person)
	fmt.Printf("person: %v\n", person)

	//结构体嵌套
	tom := Master{
		Dog{
			"虎皮",
			"花纹",
			1,
		},
		"tom",
		29,
	}
	fmt.Println(tom)

	//方法
	var per Person
	per.name = "james"
	per.eat()
	per.sleep()

	//值类型结构体和指针类型结构体
	p1 := Person{name: "jack"}
	fmt.Printf("p1: %T\n, p1.name=%v\n", p1, p1.name)
	p2 := &Person{name: "yao"}
	fmt.Printf("p2: %T\n,p2.name=%v\n", p2, p2.name)


	ren := Ren{name: "拜登"}
	fmt.Println(ren.name)
	showRen(ren)
	fmt.Println(ren.name)
	showRen1(&ren)
	fmt.Println(ren.name)

	fmt.Println("---------------------------------")
	ren.name="拜登"
	ren.showRen()
	fmt.Println(ren)
	ren.showRen1()
	fmt.Println(ren)

}
