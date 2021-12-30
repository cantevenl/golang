package main

import "fmt"

type Jisuanqi struct {
	n1 float64
	n2 float64
}

func (jisuan *Jisuanqi) Jisuan(fuhao string) {
	switch fuhao {
	case "+":
		fmt.Println(jisuan.n1 + jisuan.n2)
	case "-":
		fmt.Println(jisuan.n1 - jisuan.n2)
	case "*":
		fmt.Println(jisuan.n1 * jisuan.n2)
	case "/":
		fmt.Println(jisuan.n1 / jisuan.n2)
	default:
		fmt.Println("请输入+-*/")
	}
}

func main() {
	var jisuan Jisuanqi
	var fuhao string
	fmt.Println("第一个数：")
	fmt.Scanln(&jisuan.n1)
	fmt.Println("第二个数：")
	fmt.Scanln(&jisuan.n2)
	fmt.Println("要使用什么符号计算")
	fmt.Scanln(&fuhao)

	jisuan.Jisuan(fuhao)
}
