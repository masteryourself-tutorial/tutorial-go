package main

import "fmt"

// ============ 初始化 ============
// slice1 元素内容是 [秦始皇 牛顿 面壁者], 元素个数是 3, 容量是 3
//
// ============ 扩容并且超过 cap(), 修改数组位置不会影响 slice1 ============
// slice1 元素内容是 [秦始皇 牛顿 面壁者], 元素个数是 3, 容量是 3
// slice2 元素内容是 [始皇大帝 牛顿 面壁者 三体小说], 元素个数是 4, 容量是 6
//
// ============ 扩容未超过 cap(), 修改数组位置会影响 slice2 ============
// slice1 元素内容是 [秦始皇 牛顿 面壁者], 元素个数是 3, 容量是 3
// slice2 元素内容是 [秦始皇大帝 牛顿 面壁者 三体小说], 元素个数是 4, 容量是 6
// slice3 元素内容是 [秦始皇大帝 牛顿 面壁者 三体小说 刘慈欣], 元素个数是 5, 容量是 6
func main() {
	var slice1 = []string{"秦始皇", "牛顿", "面壁者"}
	fmt.Printf("\n ============ 初始化 ============ \n")
	fmt.Printf("slice1 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice1, len(slice1), cap(slice1))
	// 如果数组扩容不够(超过了切片的 cap), 则需要重新开辟一块内存存放扩容数据
	slice2 := append(slice1, "三体小说")
	slice2[0] = "始皇大帝"
	// 这里的 slice1、slice2 对应的数组不是同一个(因为重新开辟了一个内存存放扩容后的数组)
	fmt.Printf("\n ============ 扩容并且超过 cap(), 修改数组位置不会影响 slice1 ============ \n")
	fmt.Printf("slice1 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice2, len(slice2), cap(slice2))
	// 如果数组扩容(没超过切片的 cap)够了, 则就在原来的数组上修改, 这里会导致两边的值一起修改
	slice3 := append(slice2, "刘慈欣")
	slice3[0] = "秦始皇大帝"
	// 这里的 slice2、slice3 对应的是同一个数组
	fmt.Printf("\n ============ 扩容未超过 cap(), 修改数组位置会影响 slice2 ============ \n")
	fmt.Printf("slice1 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3 元素内容是 %v, 元素个数是 %v, 容量是 %v\n", slice3, len(slice3), cap(slice3))
}
