package main

import "fmt"

func main() {
	//方式1
	//定义map变量
	var a map[int]string
	//只声明map内存是没有分配空间
	//必须通过make函数进行初始化，才会分配空间
	a = make(map[int]string, 10) //10指的是 map可以存放10个键值对
	//将键值对存入map中：
	a[20095452] = "张三"
	a[20095497] = "李四"
	a[20097291] = "王五"
	a[20095497] = "猪六"
	a[20096699] = "张三"

	//输出集合
	fmt.Println(a)

	//方式2
	b := make(map[int]string)
	b[110] = "牛"
	b[120] = "猪"
	fmt.Println(b)

	//方式3
	c := map[int]string{
		250:  "小田",
		3600: "白眼狼",
	}
	c[28000] = "傻逼"
	fmt.Println(c)

}
