package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println("自定义错误为：", err)
		panic(err)
	}
	fmt.Println("上面除法操作完成")
}

func test() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		//抛出自定义错误
		return errors.New("除数不能为0！")
	} else { //如果除数不为0，那么正常执行就可以了
		result := num1 / num2
		fmt.Println(result)
		//如果没有错误的话返回一个零值
		return nil
	}
}
