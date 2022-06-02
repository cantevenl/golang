package main

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"unsafe"
)

func niu(a, b int) int {
	return a + b
}

func pai() {
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)
}

type user struct {
	name string
}

func main() {
	var name string = "zhu"
	age := 2
	b := true

	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", b)

	a := 100
	p := &a
	fmt.Printf("%T %v\n", p, *p)

	c := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", c)

	d := []int{1, 2, 3}
	fmt.Printf("%T\n", d)

	fmt.Printf("%T\n", niu)
	fmt.Printf("%T\n", niu(1, 2))

	var b1 bool = true
	var b2 = false
	b3 := true
	fmt.Println(b1, b2, b3)

	age1 := 10
	ok := age1 >= 18
	if ok {
		fmt.Println("你已经成年")
	} else {
		fmt.Println("你还未成年")
	}

	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("i=%v\n", i)
	}

	age2 := 18
	gender := "man"
	if age2 >= 18 && gender == "man" {
		fmt.Println("18+man")
	}

	var i8 int8
	var ui8 uint8
	fmt.Printf("%T %dB %v~%v\n ", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
	fmt.Printf("%T %dB %v~%v\n ", i8, unsafe.Sizeof(ui8), 0, math.MaxUint8)

	pai()

	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	var str1 string = "hello world"
	var html string = `
	<html>
	<head><title>hello golang</title>
	</html>
		`

	fmt.Printf("str1: %v\n", str1)
	fmt.Printf("html: %v\n", html)

	msg := ""
	age3 := "20"
	msg += name
	msg += " "
	msg += age3
	fmt.Printf("msg: %v\n", msg)

	name1 := "zhangsan"
	age4 := "27"
	msg1 := fmt.Sprintf("name:%s---age:%s", name1, age4)
	fmt.Println(msg1)

	name2 := "lisi"
	age5 := "30"
	msg2 := strings.Join([]string{name2, age5}, "+")
	fmt.Printf("msg2=%v\n", msg2)

	//这个比较理想，可以当成可变字符使用，对内存的增长也有优化，如果能预估字符串的长度，还可以用 buffer.Grow() 接口来设置 capacity
	var buffer bytes.Buffer
	buffer.WriteString("tom")
	buffer.WriteString("@")
	buffer.WriteString("80")
	fmt.Printf("buffer.String():%v \n",buffer.String())

	//str := "hello world"
	str := "abcde fghiz"
	n := 9
	m := 10
	fmt.Println(len(str))
	fmt.Println(str[m])   //获取字符串索引位置为n的原始字节
	fmt.Println(str[n:m]) //截取得字符串索引位置为 n 到 m-1 的字符串
	fmt.Println(str[n:])  //截取得字符串索引位置为 n 到 len(s)-1 的字符串
	fmt.Println(str[:m])  //截取得字符串索引位置为 0 到 m-1 的字符串

	s := "hello world!niu"
	s1 := []string{"1","a","c","2"}
	fmt.Printf("len(s): %v\n", len(s))	//获取字符串长度
	fmt.Printf("strings.Split(s, \"\"): %v\n", strings.Split(s, "!"))	//分割
	for i,v := range s {
		fmt.Printf("i=%v,\tv=%v\n",i,v)
	}
	fmt.Printf("strings.Contains(s, \"hello\"): %v\n", strings.Contains(s, "hello"))	//判断是否包含
	fmt.Printf("strings.HasPrefix(s, \"hello\"): %v\n", strings.HasPrefix(s, "hello")) //前缀判断
	fmt.Printf("strings.HasSuffix(s, \"world！\"): %v\n", strings.HasSuffix(s, "world！")) //后缀判断
	fmt.Printf("strings.Index(s, \"l\"): %v\n", strings.Index(s, "l"))  //子串出现的下标位置
	fmt.Printf("strings.LastIndex(s, \"l\"): %v\n", strings.LastIndex(s, "l"))   //子串最后出现的下标
	fmt.Printf("strings.Join(s,\"zhu*\")：%v\n------------------------------------------\n",strings.Join(s1,"*zhu"))

	u := user{"zhang"}
	//Printf 格式化输出
	fmt.Printf("% + v\n", u)     //格式化输出结构
	fmt.Printf("%#v\n", u)       //输出值的 Go 语言表示方法
	fmt.Printf("%T\n", u)        //输出值的类型的 Go 语言表示
	fmt.Printf("%t\n", false)     //输出值的 true 或 false
	fmt.Printf("%b\n", 1024)     //二进制表示
	fmt.Printf("%c\n", 97) //数值对应的 Unicode 编码字符  97---a
	fmt.Printf("%d\n", 10)       //十进制表示
	fmt.Printf("%o\n", 8)        //八进制表示
	fmt.Printf("%q\n", 22)       //转化为十六进制并附上单引号
	fmt.Printf("%x\n", 1223)     //十六进制表示，用a-f表示
	fmt.Printf("%X\n", 1223)     //十六进制表示，用A-F表示
	fmt.Printf("%U\n", 1233)     //Unicode表示
	fmt.Printf("%b\n", 12.34)    //无小数部分，两位指数的科学计数法6946802425218990p-49
	fmt.Printf("%e\n", 12.345)   //科学计数法，e表示
	fmt.Printf("%E\n", 12.34455) //科学计数法，E表示
	fmt.Printf("%f\n", 12.3456)  //有小数部分，无指数部分
	fmt.Printf("%g\n", 12.3456)  //根据实际情况采用%e或%f输出
	fmt.Printf("%G\n", 12.3456)  //根据实际情况采用%E或%f输出
	fmt.Printf("%s\n", "wqdew")  //直接输出字符串或者[]byte
	fmt.Printf("%q\n", "dedede") //双引号括起来的字符串
	fmt.Printf("%x\n", "abczxc") //每个字节用两字节十六进制表示，a-f表示
	fmt.Printf("%X\n", "asdzxc") //每个字节用两字节十六进制表示，A-F表示
	fmt.Printf("%p\n-----------------------------\n", 0x123)    //0x开头的十六进制数表示

	var a1 int
	a1 = 100
	fmt.Printf("a: %v\n", a1)
	a1 += 1 // a = a + 1
	fmt.Printf("a: %v\n", a1)
	a1 -= 1 // a = a -1
	fmt.Printf("a: %v\n", a1)
	a1 *= 2 // a = a * 2
	fmt.Printf("a: %v\n", a1)
	a1 /= 2 // a = a / 2
	fmt.Printf("a: %v\n", a1)

	var num int
	fmt.Println("输入一个数字：")
	fmt.Scan(&num)
	if num%2==0 {
		fmt.Println("是偶数\n")
	}else {
		fmt.Println("是奇数\n")
	}

	if age3:=20;age3>=18 {
		fmt.Println("成年\n")
	}else {
		fmt.Println("未成年\n")
	}


}
