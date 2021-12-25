package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听端口失败", err)
	}
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			fmt.Println("关闭端口失败", err)
		}
	}(listen)
	// 死循环监听客户端连接
	for {
		// 这一步会阻塞, 直到有客户端连接过来
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("客户端连接失败", err)
		}
		process(conn)
	}

}

func process(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("客户端连接关闭失败")
		}
	}(conn)
	for {
		buf := make([]byte, 1024)
		read, err := conn.Read(buf)
		if err != nil {
			return
		}
		fmt.Printf("客户端地址是 %v 接收到客户端信息 %v \n", conn.RemoteAddr(), string(buf[0:read]))
	}
}
