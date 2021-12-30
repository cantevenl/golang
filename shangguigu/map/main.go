package main

import "fmt"

func main() {
	//第一种使用方式
	//map声明和注意事项
	//var a map[string]string
	//fmt.Println(a)
	//在使用map前，需要先make，make的作用是给map分配数据空间
	//a = make(map[string]string, 10)
	//a["一号"] = "刘备"
	//a["二号"] = "关羽"
	//a["三号"] = "张飞"
	//fmt.Println(a)

	var a map[string]string
	a = make(map[string]string)
	a["一号"] = "牛"
	a["二号"] = "猪"
	a["三号"] = "虎"
	fmt.Println(a)

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "成都"
	fmt.Println(cities)

	//因为 no3这个key存在，因此下面是修改
	cities["no3"] = "杭州"
	fmt.Println(cities)

	//演示删除
	delete(cities, "no1")
	fmt.Println(cities)

	//如果希望一次性删除所有的key
	//1.遍历所有的key逐一删除
	//2.直接make一个新的空间
	cities = make(map[string]string)
	fmt.Println(cities)

	//第三种方式
	heroes := map[string]string{
		"hero1": "宋江",
		"hero2": "卢俊义",
		"hero3": "吴用",
	}
	fmt.Println(heroes)
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "猪"
	studentMap["stu01"]["gender"] = "男"
	studentMap["stu01"]["address"] = "上海迪士尼"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "猫"
	studentMap["stu02"]["gender"] = "女"
	studentMap["stu02"]["address"] = "上海迪士尼"

	fmt.Println(studentMap["stu01"])
	fmt.Println(studentMap["stu02"])

	val, findRes := studentMap["stu01"]
	if findRes {
		fmt.Println("找到了val=", val)
	} else {
		fmt.Printf("没有")
	}

	//使用for-range遍历map
	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v,v2=%v\n", k2, v2)
		}
		fmt.Println()
	}

	heroes = make(map[string]string)
	fmt.Println(heroes)

	studentMap = make(map[string]map[string]string)
	fmt.Println(studentMap)

	var niu map[string]string
	niu = make(map[string]string)
	niu["1"] = "1"
	niu["2"] = "2"
	fmt.Println(niu)

	var zhu []map[string]string
	zhu = make([]map[string]string, 2)
	zhu[0] = make(map[string]string)
	zhu[0]["1"] = "1"
	zhu[1] = make(map[string]string)
	zhu[1]["2"] = "2"
	zhu2 := map[string]string{
		"3": "3",
	}
	zhu = append(zhu, zhu2)
	fmt.Println(zhu)

}
