package main

import (
	"fmt"
)

func main() {
	pupil := Pupil{}
	pupil.SetId(1)
	pupil.SetScore(165)
	pupil.Show()

	senior := Senior{}
	senior.SetId(1)
	senior.SetScore(618.65)
	senior.Show()
	senior.Stu.Show()
}

// 父类接口
type Stu struct {
	id    int32
	score float64
}

// 注意这里要用指针接收, 否则值修改没有
func (s *Stu) SetId(id int32) {
	s.id = id
}

// 注意这里要用指针接收, 否则值修改没有
func (s *Stu) SetScore(score float64) {
	s.score = score
}

func (s Stu) Show() {
	fmt.Printf("展示考试成绩： 学号 = %v, score = %v \n", s.id, s.score)
}

// 小学生
type Pupil struct {
	// 继承 Stu
	Stu
}

// 重写父类方法
func (p Pupil) Show() {
	fmt.Printf("小学生考试成绩满分200, 学号 = %v, score = %v \n", p.Stu.id, p.Stu.score)
}

// 高中生
type Senior struct {
	// 继承 Stu
	Stu
}

// 重写父类方法
func (s Senior) Show() {
	fmt.Printf("高中生考试成绩满分600, 学号 = %v, score = %v \n", s.Stu.id, s.Stu.score)
}
