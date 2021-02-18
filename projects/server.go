package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net listen err: ", err)
		return
	}
	defer listen.Close()
	fmt.Println("服务器等待客户端建立连接...")
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("listen.Accept() err: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("服务器与客户端成功建立连接！！！")
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read(buf) err: ", err)
		return
	}
	conn.Write(buf)
	fmt.Println("服务器读取到的数据是：", string(buf[:n]))
}
