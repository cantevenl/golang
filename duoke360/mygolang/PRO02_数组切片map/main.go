package main

import "fmt"

func main() {
	var a [3]int
	var s [5]string
	var a1 = [...]int{1, 2, 3, 4, 5}
	var s1 = [...]string{"a", "b"}
	var b1 = [...]bool{true, false}
	//指定索引值的方式来初始化
	var a2 = [...]int{1: 100, 4: 400}
	a2[0] = 10000
	a2[3] = 333

	fmt.Printf("a: %T\n", a)
	fmt.Printf("s: %T\n", s)

	fmt.Println(a1, s1, b1)
	fmt.Println(a2)
	fmt.Println(a2[3])

	a3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < len(a3); i++ {
		fmt.Printf("a3[%d]: %v\n", i, a3[i])
	}
	fmt.Println("-----------for range------------")

	for i, v := range a3 {
		fmt.Printf("a3[%v]=%v\n", i, v)
	}

	var slice []int
	var slice1 []string

	slice = []int{1, 2, 3, 4}
	slice1 = []string{"a", "b"}

	fmt.Println(slice, slice1)

	var slice2 = []bool{true, false}
	fmt.Println(slice2)

	slice3 := []int{1, 2, 3, 4}
	fmt.Println(slice3)

	slice4 := make([]string, 5)
	slice4 = []string{"niu", "zhu"}
	fmt.Println(slice4)

	var slice5 = make([]int, 3, 5)
	fmt.Printf("len:%v,cap:%v,Type:%T\n", len(slice5), cap(slice5), slice5)

	var array = []int{1, 2, 3, 4, 5, 6, 7}
	slice6 := array[0:3]
	fmt.Println(slice6)
	slice7 := array[3:]
	slice8 := array[2:5]
	slice9 := array[:]
	fmt.Println(slice7, slice8, slice9)

	for i := 0; i < len(slice9); i++ {
		fmt.Printf("slice9[%d]：%v\n", i, slice9[i])
	}

	for i, v := range slice9 {
		fmt.Printf("slice9[%d]：%v\n", i, v)

	}

	var slice10 = []int{}
	slice10 = append(slice10, 10)
	slice10 = append(slice10, 20, 30, 40, 50, 60, 70)
	fmt.Println(slice10)

	//删除指定下标的元素
	//最后有三个.
	slice10 = append(slice10[:3], slice10[4:]...)
	fmt.Println(slice10)

	//拷贝切片
	s2 := []int{1, 2, 3}
	s3 := s2
	s3[0] = 100
	fmt.Printf("s2：%v\ns3：%v\n", s2, s3)
	fmt.Println("-----------------")

	s4 := make([]int, 3)
	copy(s4, s2)
	s4[0] = 999
	fmt.Printf("s2：%v\ns4：%v\n", s2, s4)

	m1 := make(map[string]string)
	m1["name"] = "tom"
	m1["addr"] = "chengdu"

	m2 := map[string]string{
		"name": "jack",
		"addr": "beijing",
	}

	fmt.Println(m1, m2)

	name := m1["name"]
	fmt.Println(name)

	//判断某个key是否存在
	v, ok := m2["addr"]
	if ok {
		fmt.Printf("存在,%v\n", v)
	} else {
		fmt.Println("不存在")
	}

	m3 := make(map[string]string)
	m3["name"] = "tom"
	m3["age"] = "18"
	m3["addr"] = "chengdu"

	for k, v := range m3 {
		fmt.Printf("m3[%v]=%v\n", k, v)
	}
	fmt.Println("-------------------------------")
	m4 := m3
	fmt.Printf("m4：%v\n", m4)

	m4["age"] = "50"
	fmt.Println(m4["age"], m3["age"],"改变了m3的值，说明map是地址传递")
}
