package main

import (
	"fmt"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

//Less 方法就是决定你使用什么标准进行排序
//按 Hero 的年龄从小到大排序!!
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	//修改成对 Name 排序
	//return hs[i].Name < hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	//交换
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp
	//下面的一句话等价于三句话
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {
	hero := HeroSlice{
		{"niu", 18},
		{"zhu", 3},
		{"mao", 1},
	}

	fmt.Println(hero)
	sort.Sort(hero)
	fmt.Println(hero)
}
