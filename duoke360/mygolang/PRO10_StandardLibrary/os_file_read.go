package main

import (
	"fmt"
	"os"
)

//打开文件
func openCloseFile() {
	f, err := os.Open("test2.txt")
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}(f)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}

	//根据第二个参数，可以读写或者创建
	f1, err := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 0755)
	defer func(f1 *os.File) {
		err := f1.Close()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}(f1)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f1.Name(): %v\n", f1.Name())
	}
}

// 创建文件
func createFile() {
	// 等价于：OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	f, _ := os.Create("a2.txt")
	fmt.Printf("f.Name(): %v\n", f.Name())
	// 第一个参数 目录默认：Temp 第二个参数 文件名前缀
	f2, _ := os.CreateTemp("", "temp")
	fmt.Printf("f2.Name(): %v\n", f2.Name())
}

// read
func readOps() {
	//f, _ := os.Open("a1.txt")
	//defer func(f *os.File) {
	//	err := f.Close()
	//	if err != nil {
	//		fmt.Printf("err: %v\n", err)
	//	}
	//}(f)

	//for {
	//	buf := make([]byte, 3)
	//	n, err := f.Read(buf)
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Printf("n: %v\n", n)
	//	fmt.Printf("string(buf): %v\n", string(buf))
	//}

	// 从下标3开始读5个字节
	//buf := make([]byte, 5)
	//n, err := f.ReadAt(buf, 3)
	//if err!=nil {
	//	fmt.Printf("err: %v\n", err)
	//}else {
	//	fmt.Printf("n: %v\n", n)
	//	fmt.Printf("string(buf): %v\n", string(buf))
	//}

	// 定位
	f, _ := os.Open("a1.txt")
	f.Seek(3, 0)
	buf := make([]byte, 10)
	n, _ := f.Read(buf)
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %v\n", string(buf))
	f.Close()

}

func main() {
	//openCloseFile()
	readOps()
}
