package main

import "fmt"

func test110(s string) (string, int) {
	switch s {
	case "a":
		return "a", 0
	case "b":
		return "b", 0
	case "m":
		return "", 1
	case "c":
		return "c", 0
	default:
		return "", 0
	}
}

func main() {
	_, c := test110("m")
	fmt.Println(c)

}
