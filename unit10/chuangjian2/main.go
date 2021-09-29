package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	var t *Student = &Student{"猪儿粑", 2, "男"}
	(*t).Name = "猪儿粑"
	t.Age = 2
	t.Gender = "男"
	fmt.Println(*t)
}
