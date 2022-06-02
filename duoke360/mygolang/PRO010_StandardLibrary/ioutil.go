package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func testReadall() {
	//r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	//返回一个file结构体，实现了Reader这个接口
	f ,_ := os.Open("test.txt")
	defer f.Close()
	if b1,err := ioutil.ReadAll(f) ; err != nil {
		log.Fatal(err)
	}else {
		fmt.Printf("%s",b1)
	}
}

//遍历目录
func testReadDir() {
	if file,err := ioutil.ReadDir("."); err !=nil {
		log.Fatal(err)
	}else {
		for _,v:=range file {
			fmt.Printf("v.Name():%v\n",v.Name())
		}
	}
}

func testReadFile() {
	b, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

func testWriteFile() {
	ioutil.WriteFile("a.txt",[]byte("Sioeye"),0664)
}

func testTempFile() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tmpfile.Name():%v\n",tmpfile.Name())

	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func wri() {
	err := ioutil.WriteFile("./zhu.txt", []byte("www.zhu.com"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func rea() {
	content, err := ioutil.ReadFile("./zhu.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func main() {
	//testReadall()
	//testReadDir()
	//testReadFile()
	//testWriteFile()
	//testTempFile()
	wri()
	rea()
}
