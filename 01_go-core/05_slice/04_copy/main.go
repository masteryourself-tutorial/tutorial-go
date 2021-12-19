package main

import "fmt"

func main() {
	var slice1 = []int32{1, 2, 3, 4, 5}
	var slice2 = make([]int32, 10, 10)
	// 将 slice1 值 copy 到 slice2 中
	copy(slice2, slice1)
	// slice1 = [1 2 3 4 5]
	fmt.Printf("slice1 = %v\n", slice1)
	// slice2 = [1 2 3 4 5 0 0 0 0 0]
	fmt.Printf("slice2 = %v\n", slice2)
	slice1[0] = 10
	// slice1 = [10 2 3 4 5]
	fmt.Printf("slice1 = %v\n", slice1)
	// slice2 = [1 2 3 4 5 0 0 0 0 0]
	fmt.Printf("slice2 = %v\n", slice2)
}
