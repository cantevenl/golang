package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //文件或者目录存在
		fmt.Println("存在")
		return true, nil
	} else if os.IsNotExist(err) {
		fmt.Println("不存在")
		return false, nil
	}
	return false, err
}

func main() {
	//将F:/abc.txt文件内容导入到E://kkk.txt

	//1.首先将F:/abc.txt内容读取到内存
	//2.将读取到的内容写入E://kkk.txt

	file1Path := "F:/abc.txt"
	file2Path := "E://kkk.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("read file err = ", err)
		return
	}

	err = ioutil.WriteFile(file2Path, data, 0666)
	if err != nil {
		fmt.Println("write file err = ", err)
	}

	PathExists("F:/niu.txt")

}
