package main

import "fmt"

type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

type Plane struct {
	vendor string
	model  string
}

type Car struct {
	factory, model string
}

func (h *Human) getName() string {
	return h.firstName + "," + h.lastName
}

func (p Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

func main() {
	h1 := Human{"猪", "儿"}
	p1 := &h1
	fmt.Println(*p1)
	interfaces := []IF{}
	h := new(Human)
	h.firstName = "豬"
	h.lastName = "兒吧"
	interfaces = append(interfaces, h)
	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)
	for _, f := range interfaces {
		fmt.Println(f)
		fmt.Println(f.getName())
	}
	p := Plane{}
	p.vendor = "testVendor"
	p.model = "testModel"
	fmt.Println(p.getName())
}
