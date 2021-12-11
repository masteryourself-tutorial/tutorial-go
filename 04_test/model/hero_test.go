package model

import (
	"os"
	"testing"
)

// 文件名称必须要是 xx_test.go
// 方法名称必须是 TestXxx
func TestHero_Store(t *testing.T) {
	// 定义 Hero 对象
	var hero = Hero{
		Name:  "迪迦",
		Age:   18,
		Skill: []string{"电磁脉冲", "死亡射线"},
	}
	hero.Store()
	// 校验是否有文件生成
	_, err := os.Open(fileName)
	if err != nil {
		t.Fatal("单元测试失败 \n", err)
	}
}

func TestHero_ReStore(t *testing.T) {
	var hero Hero
	hero.ReStore()
	t.Logf("hero = %v \n", hero)
}
