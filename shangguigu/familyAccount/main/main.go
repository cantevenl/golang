package main

import (
	"fmt"
	"golang/shangguigu/familyAccount/utils"
)

func main() {
	fmt.Println("这是面向对象的方式完成")
	utils.NewFamilyAccount().MainMenu()
}
