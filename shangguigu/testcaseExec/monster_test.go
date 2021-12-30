package monster

import "testing"

func TestMonster_Store(t *testing.T) {
	///先创建一个Monster实例
	monster := &Monster{
		"猪八戒",
		999,
		"睡觉",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误，希望是=%v，实际是=%v", true, res)
	}
	t.Logf("monster.Store() 测试成功")
}

func TestMonster_ReStore(t *testing.T) {
	//先创建一个Monster实例，不需要指定字段的值
	//var monster Monster
	var monster = &Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() 错误，希望是=%v，实际是=%v", true, res)
	}

	//进一步判断
	if monster.Name != "猪八戒" {
		t.Fatalf("monster.ReStore() 错误，希望是=%v，实际是=%v", "猪八戒", monster.Name)
	}

	t.Logf("monster.ReStore() 测试成功")
}
