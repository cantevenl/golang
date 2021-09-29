package main

import "fmt"

//func main() {
//	//求1-100的和，当和第一次超过300的时候，停止程序
//	var sum int = 0
//	for i:=1;i<=100;i++ {
//		sum+=i
//		fmt.Println(sum)
//		if sum >=300 {
//			//停止正在执行的这个循环,跳出循环
//			break
//		}
//	}
//	fmt.Println("到300了")
//}

//func main () {
//	for i:=1;i<=5;i++ {
//		for j:=1;j<=4;j++ {
//			fmt.Printf("i:%v,j:%v\n",i,j)
//			if j>=2 {
//				break
//			}
//		}
//	}
//}

func main() {
label1:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 4; j++ {
			fmt.Printf("i:%v,j:%v\n", i, j)
			if j >= 2 {
				break label1
			}
		}
	}
	fmt.Println("niu......")
}
