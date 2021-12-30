package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("F:/k8s.txt")
	if err != nil {
		fmt.Println(err)
	}
	//当函数退出时，要及时关闭文件
	defer file.Close() //要及时关闭file句柄，否则会有内存泄漏

	//创建一个 *Reader，是带缓冲的
	/*
		const(
		defaultBufSize = 4096 //默认的缓冲区为4096
		)
	*/

	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束一次
		if err == io.EOF {                  //io.EOF表示文件的末尾
			break
		}
		//输出内容
		fmt.Print(str)
	}
	fmt.Println("文件读取完毕。。。")
}
