package main

import "fmt"

type Niu struct {
	reduis float64
}

func (n Niu) are() float64 {
	fmt.Printf("are()方法里面n结构体的地址：%p\n", &n)
	return 3.14 * n.reduis * n.reduis
}

func (n *Niu) ares() float64 {
	fmt.Printf("ares()方法里面n结构体的指针指向的地址：%p\n", n)
	fmt.Printf("ares()方法里面n结构体的指针的地址：%p\n", &n)
	return 3.14 * (*n).reduis * (*n).reduis
}

func main() {
	//var n Niu
	//n.reduis=10
	//var n *Niu= &Niu{10}
	var n *Niu = new(Niu)
	(*n).reduis = 10
	fmt.Printf("main（）n结构体指针指向的地址：%p\n", n)
	fmt.Printf("main（）n结构体的地址：%p\n", &n)
	re := n.are()
	fmt.Println(re)
	res := n.ares()
	fmt.Println(res)
}
