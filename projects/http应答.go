package main

import (
	"fmt"
	"net"
	"os"
)

func errFunc(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		os.Exit(1)
	}
}
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8006")
	errFunc(err, "net.Dial err ")
	defer conn.Close()
	httpRequest := "GET /backbay2 HTTP/1.1\r\nHost:127.0.0.1:8006\r\n\r\n"
	conn.Write([]byte(httpRequest))
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	errFunc(err, "conn.Read err ")
	if n == 0 {
		return
	}
	fmt.Printf("|%s|", string(buf[:n]))
}
