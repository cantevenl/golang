package main

import (
	"fmt"
	"os"
)

func main() {
	//打开文件
	//概念说明：file的叫法
	//1. file叫 file对象
	//2. file叫 file指针
	//3. file叫 file文件句柄
	file, err := os.Open("F:/k8s.txt")
	if err != nil {
		fmt.Println("open file err= ", err)
	}
	//输出文件
	fmt.Printf("file=%v", file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
