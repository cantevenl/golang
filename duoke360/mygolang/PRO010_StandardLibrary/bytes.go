package main

import (
	"bytes"
	"fmt"
	"io"
)

func testTrans() {
	var i int = 100
	var b byte = 10
	b = byte(i)
	fmt.Println(b)

	var s string = "hello world"
	b1 := []byte{1, 2, 3}
	s = string(b1)
	b1 = []byte(s)
	fmt.Println(s, b1)
}

func testContains() {
	b := []byte("duoke360.com") //字符串强转为byte切片
	sublice1 := []byte("duoke360")
	sublice2 := []byte("DuoKe360")
	fmt.Println(bytes.Contains(b, sublice1)) //true
	fmt.Println(bytes.Contains(b, sublice2)) //false
}

func testCount() {
	s := []byte("hellooooooooo")
	sep1 := []byte("h")
	sep2 := []byte("l")
	sep3 := []byte("o")
	fmt.Println(bytes.Count(s, sep1)) //1
	fmt.Println(bytes.Count(s, sep2)) //2
	fmt.Println(bytes.Count(s, sep3)) //9
}

func testRepeat() {
	b := []byte("hi")
	fmt.Println(string(bytes.Repeat(b, 1))) //hi
	fmt.Println(string(bytes.Repeat(b, 3))) //hihihi
}

func testReplace() {
	s := []byte("hello,world")
	old := []byte("o")
	news := []byte("ee")
	fmt.Println(string(bytes.Replace(s, old, news, 0)))  //hello,world
	fmt.Println(string(bytes.Replace(s, old, news, 1)))  //hellee,world
	fmt.Println(string(bytes.Replace(s, old, news, 2)))  //hellee,weerld
	fmt.Println(string(bytes.Replace(s, old, news, -1))) //hellee,weerld
}

func testRunes() {
	s := []byte("你好世界")
	r := bytes.Runes(s)
	fmt.Println("转换前字符串的长度: ", len(s)) //12
	fmt.Println("转换后字符串的长度: ", len(r)) //4
}

func testJoin() {
	s2 := [][]byte{[]byte("你好"), []byte("世界")}
	sep4 := []byte(",")
	fmt.Println(string(bytes.Join(s2, sep4))) //你好,世界
	sep5 := []byte("#")
	fmt.Println(string(bytes.Join(s2, sep5))) //你好#世界
}

func testBuffer() {
	var b bytes.Buffer // 直接定义一个Buffer变量，不用初始化，可以直接使用
	fmt.Println("b:", b)
	b2 := new(bytes.Buffer) //使用New返回Buffer变量
	fmt.Printf("b2类型%T,b2=%v\n", b2, *b2)

	s := []byte{1, 2, 3}
	b3 := bytes.NewBuffer(s) //从一个[]byte切片，构造一个Buffer
	fmt.Println("b3:", b3)

	s1 := "niubi"
	b4 := bytes.NewBufferString(s1) //从一个string变量，构造一个Buffer
	fmt.Println("b4:", b4)
}

func testWrite() {
	var b bytes.Buffer
	n, _ := b.WriteString("hello k8s")
	fmt.Printf("n:%v\n", n)
	fmt.Printf("b:%v\n", string(b.Bytes()))
}

func testRead() {
	b := bytes.NewBufferString("hello istio")

	b2 := make([]byte, 2)
	for {
		n, err := b.Read(b2)
		if err == io.EOF {
			break
		}
		fmt.Println("读取多少个字节：", n)
		fmt.Printf("b2:%v\n", string(b2[0:n]))
	}
}

func testReader() {
	data := "123456789"
	//通过[]byte创建Reader
	re := bytes.NewReader([]byte(data))
	//返回未读取部分的长度
	fmt.Println("re len : ", re.Len())
	//返回底层数据总长度
	fmt.Println("re size : ", re.Size())
	fmt.Println("------------")

	buf := make([]byte, 2)
	for {
		//读取数据
		n, err := re.Read(buf)
		if err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}

	fmt.Println("------------")

	//设置偏移量，因为上面的操作已经修改了读取位置等信息
	re.Seek(0, 0)
	for {
		//一个字节一个字节的读
		b, err := re.ReadByte()
		if err != nil {
			break
		}
		fmt.Println(string(b))
	}
	fmt.Println("------------")

	re.Seek(0, 0)
	off := int64(0)
	for {
		//指定偏移量读取
		n, err := re.ReadAt(buf, off)
		if err != nil {
			break
		}
		off += int64(n)
		fmt.Println(off, string(buf[:n]))
	}
}


func main() {
	//testTrans()
	//testContains()
	//testCount()
	//testRepeat()
	//testReplace()
	//testRunes()
	//testJoin()
	//fmt.Println("-----------------------------")
	//testBuffer()
	//
	//testWrite()
	//fmt.Println("-----------------------------")
	//
	//testRead()

	testReader()
}
