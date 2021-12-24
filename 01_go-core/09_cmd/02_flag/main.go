package main

import (
	"flag"
	"fmt"
)

// Edit Configurations -> Program arguments
// -u root -p root -h 127.0.0.1 -P 3306
func main() {
	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user, "u", "root", "用户名称")
	flag.StringVar(&pwd, "p", "root", "密码")
	flag.StringVar(&host, "h", "host", "host 地址")
	flag.IntVar(&port, "P", 3306, "端口号")
	// 调用解析方法
	flag.Parse()
	fmt.Printf("user = %v, pwd = %v, host = %v, port = %v",
		user, pwd, host, port)
}
