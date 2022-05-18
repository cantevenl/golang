package main

import (
	"fmt"
	"os"
)

//创建文件
//func createFile() {
//	file, err := os.Create("test.txt")
//	if err != nil {
//		fmt.Printf("err:%v\n", err)
//	} else {
//		fmt.Printf("file.Name():%v\n", file.Name())
//	}
//}

//创建目录
func makeDir() {
	err := os.Mkdir("a", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	os.MkdirAll("a/b/c", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)

	}
}

//删除目录和文件
func remove() {
	//err:= os.Remove("a.txt")
	//if err!=nil {
	//	fmt.Printf("err: %v\n", err)
	//}

	//删除当前目录所有内容
	err := os.RemoveAll("a")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

}

// 修改工作目录
func chWd() {
	err := os.Chdir("f:/golangProject/src/golang/duoke360/mygolang/")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Println(os.Getwd())
}

//获得工作目录
func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
}

// 获得临时目录
func getTemp() {
	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

// 重命名文件
func renameFile() {
	err := os.Rename("test.txt", "test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

//写文件
func writeFile() {
	err := os.WriteFile("test2.txt",[]byte("hello golang"),os.ModePerm)
	if err!=nil {
		fmt.Printf("err: %v\n", err)
	}
}

//读文件
func readFile() {
	b, err := os.ReadFile("test2.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("b：%v\n", string(b[:]))
	}
}

func main() {
	//createFile()
	//makeDir()
	//remove()
	//chWd()
	//getWd()
	//getTemp()
	//renameFile()
	writeFile()
	readFile()

}
