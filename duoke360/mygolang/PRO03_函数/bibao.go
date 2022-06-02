package main

import "fmt"

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func counter(f func()) func() int {
	n:=0
	return func() int {
		f()
		n++
		return n
	}
}

func foo() {
	fmt.Println("call foo")
}

func makeFibGen() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = f1+f2, f2
		return f1
	}
}

func main() {
	f := add()
	r := f(10)
	fmt.Println(r)
	r = f(20)
	fmt.Println(r)
	r = f(30)
	fmt.Println(r)

	f1 := add()
	r= f1(40)
	fmt.Println(r)

	cnt := counter(foo)
	cnt()
	cnt()
	cnt()
	fmt.Println(cnt())

	fmt.Println("-------------------------------------")

	fibo:=makeFibGen()
	for i:=0;i<10;i++ {
		fmt.Println(fibo())
	}
}
