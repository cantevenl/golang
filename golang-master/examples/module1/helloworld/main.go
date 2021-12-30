package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//flag.String这个方法可以将参数传递进来
	//go run monster.go --name zhuerba
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fullString := fmt.Sprintf("Hello %s from Go\n", *name)
	fmt.Println(fullString)

	err, result := DuplicateStrig("aaa")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}

	err, result = DuplicateStrig1("bbbb")
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}

func DuplicateStrig(input string) (error, string) {
	if input == "aaa" {
		return fmt.Errorf("aaa is not allowed"), ""
	}
	return nil, input + input

}

func DuplicateStrig1(input string) (err error, result string) {
	if input == "bbb" {
		err = fmt.Errorf("bbb is error")
		return
	}
	result = input + input
	return
}
