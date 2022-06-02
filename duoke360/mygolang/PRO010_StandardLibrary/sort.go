package main

import (
	"fmt"
	"sort"
)

type NewInts []uint

func (n NewInts) Len() int {
	return len(n)
}

func (n NewInts) Less(i, j int) bool {
	fmt.Println(i, j, n[i] < n[j], n)
	return n[i] < n[j]
}

func (n NewInts) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func testFloat64(f []float64) {
	sort.Float64s(f)
	fmt.Printf("f:%v\n", f)
}

func testStrings() {
	//字符串排序，现比较高位，相同的再比较低位
	// [] string
	ls := sort.StringSlice{
		"100",
		"42",
		"41",
		"3",
		"2",
	}
	fmt.Println(ls)  //[100 42 41 3 2]
	sort.Strings(ls)
	fmt.Println(ls)  //[100 2 3 41 42]


	//字符串排序，现比较高位，相同的再比较低位
	ls = sort.StringSlice{
		"d",
		"ac",
		"c",
		"ab",
		"e",
	}
	fmt.Println(ls)  //[d ac c ab e]
	sort.Strings(ls)
	fmt.Println(ls)  //[ab ac c d e]


	//汉字排序，依次比较byte大小
	ls = sort.StringSlice{
		"啊",
		"博",
		"次",
		"得",
		"饿",
		"周",
	}
	fmt.Println(ls)  //[啊 博 次 得 饿 周]
	sort.Strings(ls)
	fmt.Println(ls)  //[博 周 啊 得 次 饿]

	for _, v := range ls{
		fmt.Println(v, []byte(v))
	}
}

func main() {
	n := []uint{1, 3, 2}
	sort.Sort(NewInts(n))
	fmt.Println(n)

	fmt.Println("----------------------------------")

	f := []float64{8.8, 6.6, 1.1, 3.14, 5.20}
	testFloat64(f)

	fmt.Println("----------------------------------")
	testStrings()

	fmt.Println("----------------------------------")

}
