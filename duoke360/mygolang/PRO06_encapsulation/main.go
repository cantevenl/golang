package main

import (
	"Mongo/PRO06_encapsulation/model"
	"fmt"
)

func main() {
	p := model.NewPerson("smith")
	p.Setinfo(18, 5000)
	fmt.Println(p)
	age, sal := p.Getinfo()
	fmt.Println(p.Name, "age=", age, "sal=", sal)
}
