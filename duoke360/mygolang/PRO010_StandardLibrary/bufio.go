package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func bufio_test1() {
	s := strings.NewReader("ABCEFG")
	str := strings.NewReader("123455")
	file, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	s1, _ := r.ReadString('\n')
	fmt.Println(s1)

	br := bufio.NewReader(s)
	b, err := br.ReadString('\n')
	if err == io.EOF {
		fmt.Println(err)
	}

	fmt.Println(b)

	br.Reset(str)
	b, _ = br.ReadString('\n')
	fmt.Println(b)
}

func bufio_test2() {
	s := strings.NewReader("ABCDEFGHIGKLMNOPQRSTUVWXYZ1234567890")
	br := bufio.NewReader(s)
	p := make([]byte, 10)

	for {
		n, err := br.Read(p)
		if err == io.EOF {
			break
		} else {
			fmt.Printf("string(p): %v\n", string(p[0:n]))
		}
	}
}

func bufio_test3() {
	s := strings.NewReader("ABCDEFG")
	br := bufio.NewReader(s)

	c, _ := br.ReadByte()
	fmt.Printf("%c\n", c)

	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)

	br.UnreadByte()
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
}

func bufio_test4() {
	s := strings.NewReader("你好，世界！")
	br := bufio.NewReader(s)

	c, size, _ := br.ReadRune()
	fmt.Printf("%c %v\n", c, size)

	c, size, _ = br.ReadRune()
	fmt.Printf("%c %v\n", c, size)

	br.UnreadRune()
	c, size, _ = br.ReadRune()
	fmt.Printf("%c %v\n", c, size)
}

func bufio_test5() {
	s := strings.NewReader("ABC\nDEF\r\nGHI\r\nGHI")
	br := bufio.NewReader(s)

	w, isPrefix, _ := br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)

	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)
}

func bufio_test6() {
	s := strings.NewReader("ABC,DEF,GHI,JKL")
	br := bufio.NewReader(s)

	w, _ := br.ReadSlice(',')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)

	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)
}

func bufio_test7() {
	s := strings.NewReader("ABCDEFNSJKFSF")
	br := bufio.NewReader(s)
	//b:=bytes.NewBuffer(make([]byte,0))

	//br.WriteTo(b)
	//fmt.Printf("%s\n",b)

	file, err := os.OpenFile("b.txt", os.O_CREATE|os.O_RDWR, 0777)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	br.WriteTo(file)

}

func bufio_test8() {
	file, err := os.OpenFile("a.txt", os.O_RDWR, 0777)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	if err != nil {
		log.Fatal(err)
	}
	w := bufio.NewWriter(file)
	str, err := w.WriteString("牛啊啊啊啊啊啊啊啊啊")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("str：%v\n", str)
	err = w.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func bufio_test9() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	_, err := bw.WriteString("1234214abc")
	if err != nil {
		log.Fatal(err)
	}

	c := bytes.NewBuffer(make([]byte,10))
	bw.Reset(c)

	bw.WriteString("勒布朗詹姆斯")
	bw.Flush()

	fmt.Println(b)
	fmt.Println(c)
}

func bufio_test10() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	fmt.Println(bw.Available()) // 4096
	fmt.Println(bw.Buffered())  // 0

	bw.WriteString("ABCDEFGHIJKLMN.com")
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)

	bw.Flush()
	fmt.Println(bw.Available())
	fmt.Println(bw.Buffered())
	fmt.Printf("%q\n", b)
}

func bufio_test11() {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	s := strings.NewReader("123")

	br := bufio.NewReader(s)
	rw := bufio.NewReadWriter(br, bw)
	p, _ := rw.ReadString('\n')

	fmt.Println(string(p))              //123
	rw.WriteString("asdf")
	rw.Flush()
	fmt.Println(b)                          //asdf
}

func bufio_test12() {
	r:=strings.NewReader("你牛啊")
	str, err := io.ReadAll(r)
	if err != nil {
		return
	}
	fmt.Println(string(str))

	s:=strings.NewReader("你聪明")
	br:=bufio.NewReader(s)
	b,_:=br.ReadString('\n')
	fmt.Println(b)
	//b := bytes.NewBuffer(make([]byte,10))
}

func wr() {
	// 参数2：打开模式，所有模式d都在上面
	// 参数3是权限控制
	// w写 r读 x执行   w  2   r  4   x  1
	file, err := os.OpenFile("./niu.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()
	// 获取writer对象
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello\n")
	}
	// 刷新缓冲区，强制写出
	writer.Flush()
}

func re() {
	file, err := os.Open("./niu.txt")
	if err != nil {
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
		fmt.Println(string(line))
	}

}

func main() {
	//bufio_test2()
	//bufio_test3()
	//bufio_test4()
	//bufio_test5()
	//bufio_test6()
	//bufio_test7()

	//bufio_test8()
	//bufio_test9()
	//bufio_test10()
	//bufio_test11()
	//bufio_test12()
	wr()
	re()

}


