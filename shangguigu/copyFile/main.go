package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//自己编写一个函数，接受两个文件路径 srcFileName dstFileName
func CopyFile(srcFileName, dstFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer srcFile.Close()
	//通过srcfile，获取到Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer dstFile.Close()
	//通过dstFile，获取到Writer
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer, reader)
}

func main() {
	//将E:/zhu.jpg 文件拷贝到F盘
	//调用CopyFile完成文件的拷贝
	srcFile := "E:/zhu.jpg"
	dstFile := "F:/mao.jpg"
	_, err := CopyFile(srcFile, dstFile)
	if err == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println("拷贝错误，err=", err)
	}

}
