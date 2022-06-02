package main

import (
	"bufio"
	"fmt"
	"strings"
)

func scan1() {
	s := strings.NewReader("ABC EFG 猪 牛 gou")
	fmt.Println(s)
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords)	//以空格作为分隔符进行分割
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
}

func scan2() {
	s := strings.NewReader("Hello 世界！")
	bs := bufio.NewScanner(s)
	//bs.Split(bufio.ScanBytes)
	bs.Split(bufio.ScanRunes)
	for bs.Scan() {
		fmt.Printf("%s ", bs.Text())
	}
}

func main() {
	scan1()

	scan2()
}
