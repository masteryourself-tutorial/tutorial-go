package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("客户端连接关闭失败")
		}
	}(conn)
	if err != nil {
		fmt.Println("连接远程服务失败", err)
	}
	_, err = conn.Write([]byte("hello tcp"))
	if err != nil {
		fmt.Println("写入数据失败", err)
		return
	}
}
