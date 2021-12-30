package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//创建一个新文件，并在文件内写入5句
	//1.打开文件 F:/abc.txt
	filePath := "F:/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
		return
	}

	//及时关闭file句柄，防止内存泄漏
	defer file.Close()

	//准备写入5句 "hello golang k8s"
	str := "hello golang k8s\n"
	//写入时，使用带缓冲的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的，因此写入的str还没有真正写到磁盘上，在调用WriteString方法时，
	//其实，内容是先写到缓存的，所以需要调用Flush()方法，将缓存的数据真正写入到文件中
	//否则，文件中会没有数据
	writer.Flush()
}
