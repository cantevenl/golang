package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func comp(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func f1(a int) {
	a = 200
	fmt.Printf("f1:%v\n", a)
}

func f2(s []int) {
	s[0] = 99
}

func f3(args ...int) {
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}
func f4(name string, age int, args ...int) {
	fmt.Printf("name: %v\n", name)
	fmt.Printf("age: %v\n", age)
	for _, v := range args {
		fmt.Printf("v: %v\n", v)
	}
}

type fun func(int, int) int

func cal(a, b int) int {
	return a * b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sayHello(name string) {
	fmt.Printf("Hello %s\n", name)
}

func test(name string, f func(string)) {
	f(name)
}

func sub(a, b int) int {
	return a - b
}

func jisuanqi(s string) func(int, int) int {
	switch s {
	case "+":
		return sum
	case "-":
		return sub
	case "*":
		return cal
	default:
		return nil

	}
}

func test1() {
	name := "tom"
	age := "18"
	f1 :=func() string {
		return name + age
	}
	msg := f1()
	fmt.Println(msg)
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)

	r = comp(4, 90)
	fmt.Println(r)

	a := 20000
	f1(a)
	fmt.Printf("main:%v", a)

	//地址传递会改变值
	s := []int{1, 2, 3}
	f2(s)
	fmt.Println(s)

	f3(1, 2, 3)
	fmt.Println("------------")
	f3(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println("------------")
	f4("tom", 20, 1, 2, 3)

	var f fun
	f = cal
	r1 := f(90, 200)
	fmt.Println(r1)

	f = max
	r2 := f(11, 88)
	fmt.Println(r2)

	test("tom", sayHello)

	add := jisuanqi("+")
	r11 := add(5, 8)
	fmt.Println(r11)

	cal := jisuanqi("*")
	r22 := cal(99, 77)
	fmt.Println(r22)

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}(5, 8)
	r33 := max
	fmt.Println(r33)

	test1()
}
