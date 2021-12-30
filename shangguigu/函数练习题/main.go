package main

import "fmt"

//days函数用于求出相差的天数
func days(year, month, day int) int {
	var count int
	for i := 1990; i <= year-1; i++ {
		if i%4 == 0 && i%100 != 0 || i%400 == 0 {
			count += 366
		} else {
			count += 365
		}
	}
	for j := 1; j <= month; j++ {
		switch j {
		case 1, 3, 5, 7, 8, 10:
			count += 31
		case 4, 6, 9, 11:
			count += 30
		case 2:
			if year%4 == 0 && year%100 != 0 || year%400 == 0 {
				count += 29
			} else {
				count += 28
			}

		}
	}
	count += day
	return count
}

//judge函数用来判断是打鱼还是晒网
func judge(count int) {
	fmt.Printf("距离1990年1月1日一共%v天\n", count)
	if count%5 == 0 || count%5 == 4 {
		fmt.Println("晒网")
	} else {
		fmt.Println("打鱼")
	}
}

func main() {
	var year, month, day int
	fmt.Println("请输入年份：")
	fmt.Scanln(&year)
	fmt.Println("请输入月份：")
	fmt.Scanln(&month)
	fmt.Println("请输入日期：")
	fmt.Scanln(&day)
	judge(days(year, month, day))
}
