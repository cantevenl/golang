package main

import "fmt"

func main() {
	arr := [5]string{"I", "am", "stupid", "and", "weak"}
	for index, value := range arr {
		fmt.Printf("下标为：%+v，下标对应的值为：%+v\n", index, value)
		//if value == "stupid"{
		//	arr[index] = "smart"
		//}else if value == "weak"{
		//	arr[index] = "strong"
		//}
		switch {
		case arr[index] == "stupid":
			arr[index] = "smart"
		case arr[index] == "weak":
			arr[index] = "strong"
		default:
		}

	}
	fmt.Println(arr)
}
