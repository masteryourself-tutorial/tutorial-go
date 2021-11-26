package main

import(
	"fmt"
)

func main()  {

	demo1()

	fmt.Println("======我是分隔符======")

	demo2()

	fmt.Println("======我是分隔符======")

	demo3()
}

func demo1()  {
	sum := 0
	for i := 1; i <= 100; i++ {
		if i % 9 ==0 {
			fmt.Println("能被 9 整除", i)
			sum +=i
		}
	}
	fmt.Printf("1~100 能被 9 整除个数之和是 %v \n", sum)
}

func demo2()  {
	for i := 0; i <= 6; i++ {
		fmt.Printf("%v + %v = %v \n", i, 6 - i, 6)
	}
}

func demo3()  {
	count := 1
	for {
		fmt.Println("hello world")
		count++
		if count > 10 {
			break
		}
	}
}