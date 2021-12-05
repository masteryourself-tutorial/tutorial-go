package main

func main() {
	// 打开文件
	f, err := os.Open("/Users/ruanrenzhao/go_project/tutorial-go/README.md")
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	fmt.Printf("f 的类型是 %T, 值是 %v \n", f, f)
	// 关闭文件
	if err = f.Close(); err != nil {
		fmt.Println("文件关闭失败", err)
	}
}
