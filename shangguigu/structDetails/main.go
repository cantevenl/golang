package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftup, rightDown Point
}

type Rect2 struct {
	leftup, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}

	//r1有4个int，在内存中是连续分布的
	//打印地址
	fmt.Printf("r1.leftUp.x 的地址=%p r1.leftUp.y 的地址=%p r1.rightDown.x 的地址=%p r1.rightDown.x 的地址=%p\n",
		&r1.leftup.x, &r1.leftup.y, &r1.rightDown.x, &r1.rightDown.y)

	//r2有两个 *Point类型，这两个*Point类型的本身地址也是连续的，
	//但是它们指向的地址不一定是连续的
	r2 := Rect2{&Point{1, 2}, &Point{3, 4}}
	fmt.Printf("r2.leftUp本身的地址=%p r2.rightDown.x 本身的地址=%p \n",
		&r2.leftup, &r2.rightDown)

	//他们指向的地址b不一定是连续的。。。这个要看内存当时是怎么分配的
	fmt.Printf("r2.leftUp指向的地址=%p r2.rightDown.x 指向的地址=%p \n",
		r2.leftup, r2.rightDown)

}
