package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("ResolveUDPAddr err : ", err)
		return
	}
	fmt.Println("服务器结构创建成功")
	udpConn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("ListenUDP err: ", err)
		return
	}

	defer udpConn.Close()
	fmt.Println("服务器通信Socket创建成功")
	buf := make([]byte, 4096)
	n, cltAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("ReadFromUDP err: ", err)
		return
	}
	fmt.Printf("服务器读到 %v 的数据：%s\n", cltAddr, string(buf[:n]))
	daytime := time.Now().String()
	_, err = udpConn.WriteToUDP([]byte(daytime), cltAddr)
	if err != nil {
		fmt.Println("WriteToUDP err: ", err)
		return
	}
}
