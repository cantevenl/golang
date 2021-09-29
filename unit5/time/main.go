package main

import (
	"fmt"
	"time"
)

func main() {
	//时间和日期的函数，需要导入time包，所以你获取当前时间，就要调用函数Now函数
	now := time.Now()
	fmt.Printf("now的类型是：%T,now的值是：%v\n", now, now)
	//Now()函数返回值是一个结构体，类型是time.Time

	//调用结构体中的方法：
	fmt.Printf("年：%v \n", now.Year())
	fmt.Printf("月：%v \n", now.Month())      //September
	fmt.Printf("月：%v \n", int(now.Month())) //9
	fmt.Printf("日：%v \n", now.Day())
	fmt.Printf("时：%v \n", now.Hour())
	fmt.Printf("分：%v \n", now.Minute())
	fmt.Printf("秒：%v \n", now.Second())

	fmt.Println("--------------------------------------------")
	//Printf将字符串直接输出
	fmt.Printf("当前年月日： %d-%d-%d 时分秒: %d:%d:%d  \n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	//Sprintf可以得到这个字符串，以便后续的使用
	datestr := fmt.Sprintf("当前年月日： %d-%d-%d 时分秒: %d:%d:%d  \n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	fmt.Println(datestr)

	//这个参数字符串的各个数字必须是固定的，必须这样写 必须是2006/01/02 15/04/05这个时间
	datestr2 := now.Format("2006/01/02 15/04/05")
	fmt.Println(datestr2)

	//选择任意的组合都是可以的
	datestr3 := now.Format("01-02 04:05")
	fmt.Println(datestr3)
}
