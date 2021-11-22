package main

import "fmt"

type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.len * box.width * box.height
}

func main() {
	var box = Box{
		1.1,
		2.0,
		3.0,
	}
	volumn := box.getVolumn()
	fmt.Printf("体积=%.2f", volumn)
}
