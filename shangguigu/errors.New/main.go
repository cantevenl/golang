package main

import (
	"errors"
	"fmt"
)

//函数去读取一个配置文件，如果文件名传入不正确，返回一个自定义错误

func readConf(name string) (err error) {
	if name == "config.yaml" {
		return nil
	} else {
		//返回一个自定义错误
		return errors.New("读取文件错误")
	}
}

func test() {
	err := readConf("config2.yaml")
	if err != nil {
		//如果读取文件发生错误，就输出这个错误，并终止程序
		panic(err)
	}
	fmt.Println("test()继续执行")
}

func main() {
	test()
	fmt.Println("下面代码继续执行")
}
