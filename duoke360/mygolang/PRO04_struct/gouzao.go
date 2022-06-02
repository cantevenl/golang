package main

import "fmt"

type Person1 struct {
	name string
	age  int
}

func NewPerson(name string, age int) (*Person1, error) {
	if name == "" {
		return nil, fmt.Errorf("name 不能为空")
	}
	if age < 0 {
		return nil, fmt.Errorf("age 不能为0")
	}
	return &Person1{name, age}, nil
}

func main() {
	per, err := NewPerson("tom", -20)
	if err == nil {
		fmt.Println(per)
	}else {
		fmt.Println(err)
	}
}
