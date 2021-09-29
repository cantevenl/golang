package main

import "fmt"

type Student struct {
	Age int
}

type Person struct {
	Age int
}

func main() {
	var s Student = Student{10}
	var p Person = Person{12}
	s = Student(p)
	fmt.Println(s)
	fmt.Println(p)
}
