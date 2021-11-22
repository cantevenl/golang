package main

import "fmt"

type Circle struct {
	redius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.redius * c.redius
}

//为了提高效率，通常我们方法和结构体的指针类型绑定
func (c *Circle) area2() float64 {
	fmt.Printf("c是 *Circle 指向的地址是%p\n", c)
	//因为c是指针，因此我们标准的访问其字段的方式是 (*c).redius
	return 3.14 * c.redius * c.redius
	//	等价于 return 3.14 * (*c).redius * (*c).redius
}

func main() {
	//var c Circle
	//c.redius = 4.0
	//res := c.area()
	//fmt.Println("⚪的面积是：", res)

	var c Circle
	fmt.Printf("main c 结构体变量地址=%p\n", &c)
	c.redius = 5.0
	res2 := (&c).area2()
	//编译器底层做了优化	(&c).area2() 等价 c.area2()
	//因为编译器会自动地给我们加上 &c
	fmt.Println("面鸡=", res2)
	c.redius = 10.0
	res3 := c.area2()
	fmt.Println("面积=", res3)
}
