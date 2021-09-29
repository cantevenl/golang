package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		if i == 27 {
			return
		}
	}
}
