package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//1.声明一个Hero结构体
type Hero struct {
	Name string
	Age  int
}

//2,声明一个Hero结构体的切片类型
type HeroSlice []Hero

//3.实现Interface接口，三个方法
func (hs HeroSlice) Len() int {
	return len(hs)
}

//下标为i的这个元素的值小于下标为j的这个元素的值，如果这个为真就是升序排列
//Less这方法就是决定你使用什么标准进行排序
//1.按Hero的年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	//修改成对名字排序
	//return hs[i].Name < hs[j].Name
}

//交换
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

type Student struct {
	Name  string
	Age   int
	Score int
}

type StudentSlice []Student

func (s StudentSlice) Len() int {
	return len(s)
}

func (s StudentSlice) Less(i, j int) bool {
	return s[i].Score > s[j].Score
}

func (s StudentSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	//先定义一个切片
	var intSlice = []int{-1, 0, 90, 10, 7}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//对一个结构体切片进行排序
	//定义一个切片
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(1000),
		}
		//将hero append到heroes切片
		heroes = append(heroes, hero)
	}
	//排序前的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}

	//调用sort.Sort
	sort.Sort(heroes)
	fmt.Println("-------------排序后------------")
	//排序后的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}

	fmt.Println("-------------------学生成绩-----------------------")
	//定义学生切片
	var students StudentSlice
	for i := 0; i < 8; i++ {
		student := Student{
			Name:  fmt.Sprintf("学生%d", rand.Intn(100)),
			Age:   rand.Intn(20),
			Score: rand.Intn(150),
		}
		students = append(students, student)
	}
	sort.Sort(students)
	for _, v := range students {
		fmt.Println(v)
	}
}
