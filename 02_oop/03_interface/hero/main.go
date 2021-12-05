package hero

// 定义一个英雄结构体
type Hero struct {
	Id   int32
	Name string
}

// 让 HeroSlice 实现 sort.Interface
type HeroSlice []Hero

func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Id < hs[j].Id
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
