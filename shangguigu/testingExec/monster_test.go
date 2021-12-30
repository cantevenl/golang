package main

import "testing"

//编写测试用例,用于测试Store()这个方法
func TestMonster_Store(t *testing.T) {
	//先创建一个结构体实例
	monster := &Monster{
		"猪八戒",
		999,
		"睡瞌睡",
	}
	res := monster.Store("F:/666.txt")
	if res != nil {
		t.Fatalf("monster.Store() 错误，希望是=%v，实际是=%v", nil, res)
	}
	t.Logf("monster.Store() 测试成功")
}

func TestMonster_ReStore(t *testing.T) {
	//先创建一个Monster实例，不需要指定字段的值
	var monster Monster
	res := monster.ReStore("F:/666.txt")
	if res != nil {
		t.Fatalf("monster.ReStore() 错误，希望是=%v，实际是=%v", nil, res)
	}
	t.Logf("monster.ReStore() 测试成功")

}
