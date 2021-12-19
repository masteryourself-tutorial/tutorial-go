package main

import (
	"fmt"
	"go-core/03_func/06_init/model"
)

func init() {
	fmt.Println("main init()")
}

// Stu.init() 只会初始化一次, init() 都是在 main() 执行之前初始化
func main() {
	// stu init()
	// main init()
	// main main()
	// &{} &{}
	var stu1 = new(model.Stu)
	var stu2 = new(model.Stu)
	fmt.Println("main main()")
	fmt.Println(stu1, stu2)
}
