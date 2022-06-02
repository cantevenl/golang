package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func testCopy() {
	//第一个文件拷贝到第二个文件
	r := strings.NewReader("hello world kubernetes golang istio")
	file, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, r)
	if err != nil {
		log.Fatal(err)
	}

	//written, err := io.Copy(os.Stdout, r)
	//if err != nil {
	//	log.Fatal(err)
	//}else {
	//	fmt.Printf("\nlong: %v\n",written)
	//}
}

//带有缓冲区的拷贝
func copyBuffer() {
	r1 := strings.NewReader("hello world kubernetes\n")
	r2 := strings.NewReader("hello world istio\n")
	buf := make([]byte, 8)

	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}

}

//限制
func limitRead() {
	r := strings.NewReader("hello golang\n")
	lr := io.LimitReader(r, 5)

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

//多个reader
func multiRead() {
	r1 := strings.NewReader("\nfirst reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}

//管道
func pipe() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

}

func main() {
	//r := strings.NewReader("hello world")
	//buf := make([]byte, 20)
	//r.Read(buf)
	//fmt.Printf("string(buf): %v\n", string(buf))

	testCopy()
	copyBuffer()
	limitRead()
	multiRead()
	pipe()
}
