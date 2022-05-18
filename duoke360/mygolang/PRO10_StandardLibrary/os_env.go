package main

import (
	"fmt"
	"os"
)

func main() {
	// 获得所有环境变量
	//fmt.Printf("os.Environ(): %v\n", os.Environ())

	s := os.Getenv("GOPATH")
	fmt.Printf("s: %v\n", s)
	r := os.Getenv("JAVA_HOME")
	fmt.Println("JAVA_HOME:", r)

	j,b:=os.LookupEnv("GOPATH")
	if b{
		fmt.Printf("j：%v\n",j)
	}

	os.Setenv("Zhu","MiaoMiao")
	fmt.Println(os.Getenv("Zhu"))

	//替换
	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

}
