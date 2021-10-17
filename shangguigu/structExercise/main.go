package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p1 Person
	p1.Age = 10
	p1.Name = "小明"
	var p2 *Person = &p1

	fmt.Println((*p2).Age)
	fmt.Println(p2.Age)
	p2.Name = "小张"
	fmt.Println(p2.Name, p1.Name)
	fmt.Println((*p2).Name, p1.Name)

	fmt.Printf("p1的地址是%p\n", &p1)
	fmt.Printf("p2的地址是%p,p2的值是%p", &p2, p2)
}
