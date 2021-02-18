package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8004")
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}
	defer listen.Close()
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("listen.Accept() err: ", err)
		return
	}
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}
	fileName := string(buf[:n])
	conn.Write([]byte("ok"))
	recvFile(conn, fileName)
}

func recvFile(conn net.Conn, name string) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println("os.Create err: ", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 4096)

	for {
		n, _ := conn.Read(buf)
		if n == 0 {
			fmt.Println("接收完毕！")
			return
		}
		f.Write(buf[:n])
	}
}
