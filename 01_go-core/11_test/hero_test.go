package main

import (
	"os"
	"testing"
)

// 方法名必须以 Test 开头
func TestHeroStore(t *testing.T) {
	// 定义 Hero 对象
	var hero = Hero{
		Name:  "迪迦",
		Age:   18,
		Skill: []string{"电磁脉冲", "死亡射线"},
	}
	hero.Store()
	// 校验是否有文件生成
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		t.Fatal("单元测试失败 \n", err)
		return
	}
}

func TestHeroReStore(t *testing.T) {
	var hero Hero
	hero.ReStore()
	t.Logf("hero = %v \n", hero)
}
