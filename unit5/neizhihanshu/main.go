package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//统计字符串的长度，按字节进行统计
	str := "golang你好" //在golang中，汉字是utf8，一个汉字3个字节
	fmt.Println(len(str))

	//对字符串进行遍历：
	//方式一：利用键值循环：for-range
	for i, v := range str {
		fmt.Printf("下标为：%d,具体的值为：%c\n", i, v)
	}

	//方式二：利用r:=[]rune(str)
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Println(r[i])
		fmt.Printf("%c\n", r[i])
	}

	//字符串转整数
	num1, _ := strconv.Atoi("666")
	fmt.Printf("num1的类型是：%T,num1=%v\n", num1, num1)

	//整数转字符串
	str1 := strconv.Itoa(888)
	fmt.Printf("str1的类型是:%T,str1=%v\n", str1, str1)

	//统计一个字符串有几个指定的子串
	count := strings.Count("kubernetesandgolangandistio", "and")
	fmt.Println(count)

	//不区分大小写的字符串比较
	flag := strings.EqualFold("hello", "HELLO")
	fmt.Println(flag)

	//区分大小写进行字符串比较
	fmt.Println("hello" == "Hello")

	//返回子串在字符串第一次出现的索引值，如果没有返回-1
	fmt.Println(strings.Index("kubernetesandgolangandistio", "et"))

	//字符串的替换
	fmt.Println(strings.Replace("goandjavagogo", "go", "kubernetes", -1))

	//按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组
	arr := strings.Split("gp-pyhton-java", "-")
	fmt.Println(arr)

	//字符串字母大小写进行转换
	fmt.Println(strings.ToLower("GO"))
	fmt.Println(strings.ToUpper("go"))

	//字符串左右两边空格去掉
	fmt.Println(strings.TrimSpace("       go   and   kubernetes   istio"))

	//将字符串左右两边指定的字符去掉
	fmt.Println(strings.Trim("~golong~", "~"))

	//将字符串左边指定的字符去掉，将字符串右边指定的字符去掉
	fmt.Println(strings.TrimLeft("~golong~", "~"))
	fmt.Println(strings.TrimRight("~golong~", "~"))

	//判断字符串是否以指定的字符开头，是否指定的字符结尾
	fmt.Println(strings.HasPrefix("http://mgt.xiaokakj.com/login", "http"))
	fmt.Println(strings.HasSuffix("http://mgt.xiaokakj.com/login", "logi"))

}
