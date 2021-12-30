package main

import "fmt"

type Cat struct {
	Name  string
	Age   int
	Score [2]float64
	High  *int
	Hobby []string
	Addr  map[string]string
}

func main() {
	var cat Cat
	fmt.Println(cat) //{ 0 [0 0 0 0 0] <nil> [] map[]}
	cat.Name = "喵喵"
	cat.Age = 1
	cat.Score = [...]float64{100.0, 88.88}
	high := 30
	cat.High = new(int)
	cat.High = &high
	cat.Hobby = make([]string, 10)
	cat.Hobby = []string{"打游戏，看书，学golang"}
	cat.Addr = make(map[string]string)
	cat.Addr["家庭住址"] = "成都市锦江区庆云街"

	fmt.Println(cat)

	cat1 := cat //结构体是值类型，默认值拷贝
	cat1.Name = "小老几"
	fmt.Println(cat1)

}
