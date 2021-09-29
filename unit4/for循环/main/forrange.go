package main

import "fmt"

func main() {
	var str string = "hello golang你好"
	//方式一：普通for循环：按照字节进行遍历输出的
	//for i := 0; i < neizhihanshu2(str); i++ {
	//	fmt.Printf("%c\n", str[i])
	//}

	//方式二：for range
	for i, value := range str {
		fmt.Printf("索引为:%d,具体的值:%c\n", i, value)
	}

	//对str进行遍历，遍历的每个结果的索引值被i接收，每个结果的具体数值被value接收
}
