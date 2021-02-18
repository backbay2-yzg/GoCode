package main

//
//import (
//	"fmt"
//	"net"
//	"strings"
//)
//
//func main() {
//	listen, err := net.Listen("tcp", "127.0.0.1:8001")
//	if err != nil {
//		fmt.Println("listen err: ", err)
//		return
//	}
//	defer listen.Close()
//
//	for {
//		fmt.Println("服务器等待客户端连接中")
//		conn, err := listen.Accept()
//		if err != nil {
//			fmt.Println("listen.accept err: ", err)
//			return
//		}
//
//		go HandlerConnection(conn)
//	}
//}
//
//func HandlerConnection(conn net.Conn) {
//	defer conn.Close()
//	addr := conn.RemoteAddr()
//	fmt.Println(addr, "客户端连接成功！")
//	buf := make([]byte, 4096)
//	for {
//		n, err := conn.Read(buf)
//		if err != nil {
//			fmt.Println("conn.read err: ", err)
//			return
//		}
//		fmt.Println("服务器读到数据：",addr, string(buf[:n]))
//
//		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
//	}
//
//}
