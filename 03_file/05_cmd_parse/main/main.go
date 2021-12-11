package main

import (
	"flag"
	"fmt"
	"os"
)

// Edit Configurations -> Program arguments
// -u root -p root -h 127.0.0.1 -P 3306
func main() {
	// 使用 os.Args 自己解析
	fmt.Printf("\n==================== 分隔符: parseByOsArgs ====================\n")
	parseByOsArgs()
	// 使用 flag 解析
	fmt.Printf("\n==================== 分隔符: parseByFlag ====================\n")
	parseByFlag()
}

func parseByOsArgs() {
	args := os.Args
	fmt.Printf("输入参数是 %v \n", args)
}

func parseByFlag() {
	var user string
	var pwd string
	var host string
	var port int
	flag.StringVar(&user, "u", "root", "用户名称")
	flag.StringVar(&pwd, "p", "root", "密码")
	flag.StringVar(&host, "h", "host", "host 地址")
	flag.IntVar(&port, "P", 3306, "端口号")
	fmt.Printf("user = %v, pwd = %v, host = %v, port = %v",
		user, pwd, host, port)
}
