package main

import "fmt"

func main() {
	var score int
	fmt.Println("请输入你的分数：")
	fmt.Scanln(&score)
	switch score / 10 {
	case 10, 11, 12:
		fmt.Println("你是A级")
	case 9:
		fmt.Println("你是B级")
	case 8:
		fmt.Println("你是C级")
		fallthrough
	case 7:
		fmt.Println("你是D级")
	case 6:
		fmt.Println("你是E级")
	case 5:
		fmt.Println("你是E级")
	case 4:
		fmt.Println("你是E级")
	case 3:
		fmt.Println("你是E级")
	default:
		fmt.Println("你是煞笔")
		//default分支是用来”兜底“的一个分支，其他case分支都不走的情况下就会走default分支
		//default分支可以放在任意位置上，不一定非要放在最后
	}

	var a int32 = 1
	switch {
	case a == 1:
		fmt.Println("aaa")
	case a == 2:
		fmt.Println("bbb")
	}

	switch b := 7; {
	case b > 6:
		fmt.Println("大于6")
	case b < 6:
		fmt.Println("不大于6")
	}
}
