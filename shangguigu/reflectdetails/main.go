package main

import (
	"fmt"
	"reflect"
)

//通过反射，修改 num int的值
//修改Student的值

type Student struct {
	Name string
	Age  int
}

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	//看看 rVal的Kind是
	fmt.Printf("int类型----rVal=%v\trVal的kind=%v\n", rVal, rVal.Kind())
	//得到rVal这个指针具体指向那个值，然后改变
	//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的value封装
	rVal.Elem().SetInt(20)
}

func main() {
	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num)

	//你可以这样理解rVal.Elem()
	num2 := 9
	var ptr *int = &num
	num3 := *ptr

	fmt.Println(num2, num3)

	stu := Student{"猪", 3}
	stuName := reflect.ValueOf(&stu.Name)
	stuAge := reflect.ValueOf(&stu.Age)
	stuName.Elem().SetString("猫")
	stuAge.Elem().SetInt(1)
	fmt.Println(stu)

}
