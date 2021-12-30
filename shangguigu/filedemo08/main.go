package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//定义一个结构体，保存统计结果
type CharCount struct {
	ChCount    int
	NumCount   int
	SpaceCount int
	OtherCount int
}

func main() {
	//打开一个文件，创建一个Reader
	//每读取一行，就去统计改行有多少个 英文、数字、空格和其他字符
	//然后将结果保存到一个结构体
	filename := "E:/abc.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	defer file.Close()
	//定义个CharCount实例
	var count CharCount
	//创建一个Reader
	reader := bufio.NewReader(file)

	//开始循环的读取fileName的内容
	for {
		str, err := reader.ReadString('\n')

		//遍历 str，进行统计
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透处理
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v == ' ' || v == '\t':
				count.SpaceCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			default:
				count.OtherCount++

			}
		}
		fmt.Println(str)
		if err == io.EOF { //读到文件末尾就退出
			break
		}

	}
	//输出统计的结果
	fmt.Printf("英文的个数为=%v,数字的个数为=%v,空格的个数为=%v,其它字符个数=%v", count.ChCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
