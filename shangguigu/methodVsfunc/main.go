package main

import "fmt"

type Person struct {
	Name string
}

//函数
func test01(p Person) {
	fmt.Println(p.Name)
}

func test02(p *Person) {
	fmt.Println(p.Name)
}

//对于方法（如struct的方法），接收者为值类型时，可以直接用指针类型的变量调用方法，反过来同样也可以
func (p Person) test03() {
	p.Name = "猫"
	fmt.Println("test03()=", p.Name) //猫
}

func (p *Person) test04() {
	p.Name = "狗"
	fmt.Println("test03()=", p.Name) //狗
}

func main() {
	p := Person{"猪"}
	test01(p)
	test02(&p)

	p.test03()
	fmt.Println("main() p.name=", p.Name) //猪

	(&p).test03() //从形式上看是传入了地址，但是本质上仍然是值拷贝

	fmt.Println("main() p.name=", p.Name) //猪

	(&p).test04()

	fmt.Println("main() p.name=", p.Name) //狗
	p.test04()                            //等价 (&p).test04()，从形式上看是传入了值类型，但是本质上仍然是地址拷贝

}
