package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("格式为：go run xxxxx.go 文件绝对路径")
		return
	}
	filePath := list[1]
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("os.Stat err: ", err)
		return
	}
	fileName := fileInfo.Name()
	fmt.Println(fileInfo.Name(), fileInfo.Size())
	conn, err := net.Dial("tcp", "127.0.0.1:8004")
	if err != nil {
		fmt.Println("net dial err: ", err)
		return
	}
	defer conn.Close()
	//err := conn.Write([]byte(fileName))
	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return
	}
	buf := make([]byte, 16)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}
	if "ok" == string(buf[:n]) {
		sendFile(conn, filePath)

	}
}

func sendFile(conn net.Conn, path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err: ", err)
		return
	}
	buf := make([]byte, 4096)

	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("发送文件完成。")
			} else {
				fmt.Println("f.Read err: ", err)
			}
			//3238805 3201769
			return
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write err: ", err)
			return
		}
	}
}
