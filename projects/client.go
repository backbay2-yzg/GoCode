package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8006")
	if err != nil {
		fmt.Println("net dial err: ", err)
		return
	}
	defer conn.Close()
	conn.Write([]byte("Are you ready ?"))
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn read err: ", err)
		return
	}
	fmt.Println("服务器回发：", string(buf[:n]))

}
