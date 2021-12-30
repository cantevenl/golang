package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//使用ioutil.ReadFile一次性将文件读取到位
	file := "F:/k8s.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
	}
	//读取都的内容显示到终端
	fmt.Println(string(content)) //[]byte
	//我们没有显式的Open文件，因此也不需要显式的Close文件
	//因为，文件的Open和Close被封装到ReadFile函数内部

}
