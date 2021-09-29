package main

import "fmt"

func main() {
	//定义字符类型的数据
	var c1 byte = 'a'
	fmt.Println(c1)
	var c2 byte = '6'
	fmt.Println(c2)
	var c3 byte = '('
	fmt.Println(c3 + 20)
	//字符类型，本质就是一个整数，也可以直接参加运算，输出字符的时候，会将对应的码值做一个输出
	//字母，数字，标点等字符，底层是按照ASCII进行存储的

	var c4 int = '中'
	fmt.Println(c4)
	//汉字字符，底层对应的是Unicode码值
	//对应的码值为20013，byte类型溢出，能存储的范围：可以用int
	//总结：Golang的字符对应的使用的是UTF-8编码（Unicode是对应的字符集，UTF-8是Unicode的其中的一种编码方案）

	var c5 byte = 'A'
	//想显示对应的字符，必须采用格式化输出
	fmt.Printf("c5对应的具体的字符为：%c\n", c5)

	var c6 int = '中'
	fmt.Printf("c6对应的字符为：%c\n", c6)

	//练习转义字符
	// \n换行
	fmt.Println("aaa\nbbb")
	// \b退格，王倩退一格
	fmt.Println("aaa\bbbb")
	// \r光标回到本行的开头,后续输入就会替换原有的字符
	fmt.Println("aaaaa\rbcd")
	// \t制表符
	fmt.Println("aaa\tbbb")
	// \" 转义
	fmt.Println("aaaaa\"aaaa")

}
