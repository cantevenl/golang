package main

import "fmt"

func testAppend() {
	slice := []int{1, 2, 3}
	i := append(slice, 100)
	fmt.Println(i)

	slice1:= []int{10,20,30}
	i2 := append(slice,slice1...)
	fmt.Println(i2)
}

func testPanic() {
	defer fmt.Println("panic 异常后执行...")
	//panic("panic 错误...")
	fmt.Println("end...")
}

func testnew() {
	b:=new(bool)
	fmt.Printf("b类型%T，b的值%v",b,*b)
}

func testmake() {
	var p11 *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
	fmt.Printf("p11:%v\n",p11)
	var v11  []int = make([]int, 10) // the slice v now refers to a new array of 100 ints
	fmt.Printf("v11:%v\n",v11)
	// Unnecessarily complex:这种做法实在是很蛋疼

	var p *[]int = new([]int)
	*p = make([]int, 10, 100)

	// Idiomatic:习惯的做法
	v12 := make([]int, 10)
	fmt.Printf("v12:%v\n",v12)

}


func main() {
	testnew()

	testmake()
}
