package main

import "fmt"

func f11() {
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	default:
		fmt.Println("一般")
	}
}

func f12() {
	day := 3
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("工作日")
	case 6, 7:
		fmt.Println("休息日")
	}
}

func f13() {
	a:=200
	switch a {
	case 100:
		fmt.Println("100")
	case 200:
		fmt.Println("200")
		//break
		fallthrough
	case 300:
		fmt.Println("300")
	default:
		fmt.Println("niu")
	}
}

func main() {
	f11()
	f12()
	f13()
}
